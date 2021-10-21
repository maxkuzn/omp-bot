package railwaystation

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/travel"
)

type CallbackListData struct {
	LastID uint64 `json:"last_id"`
	GoNext bool   `json:"next"`
}

func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("railwaystation.Commander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	var stations []travel.RailwayStation
	if parsedData.GoNext {
		stations, err = c.service.List(parsedData.LastID+1, listLimit)
		if err != nil {
			reply(c.bot, callback.Message.Chat.ID, fmt.Sprintf("An error occured: %s", err))
			return
		}
	} else {
		stations, err = c.service.ListUntil(parsedData.LastID, listLimit)
		if err != nil {
			reply(c.bot, callback.Message.Chat.ID, fmt.Sprintf("An error occured: %s", err))
			return
		}
	}

	var text string
	if len(stations) != 0 {
		text = "List of stations:\n"
		for _, s := range stations {
			text += fmt.Sprintln(s)
		}
	} else {
		text = "There are no more stations"
	}

	callbackDataPrev := CallbackListData{
		GoNext: false,
	}
	if len(stations) != 0 {
		callbackDataPrev.LastID = stations[0].ID
	} else if parsedData.GoNext {
		callbackDataPrev.LastID = parsedData.LastID + 1
	}
	serializedDataPrev, err := json.Marshal(callbackDataPrev)
	if err != nil {
		panic(err)
	}
	callbackPathPrev := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "railway_station",
		CallbackName: "list",
		CallbackData: string(serializedDataPrev),
	}

	callbackDataNext := CallbackListData{
		GoNext: true,
	}
	if len(stations) != 0 {
		callbackDataNext.LastID = stations[len(stations)-1].ID
	} else if parsedData.GoNext {
		callbackDataNext.LastID = parsedData.LastID - 1
	}
	serializedDataNext, err := json.Marshal(callbackDataNext)
	if err != nil {
		panic(err)
	}
	callbackPathNext := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "railway_station",
		CallbackName: "list",
		CallbackData: string(serializedDataNext),
	}

	var replyMarkup tgbotapi.InlineKeyboardMarkup
	switch {
	case len(stations) != 0:
		replyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPathPrev.String()),
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			),
		)
	case parsedData.GoNext:
		replyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPathPrev.String()),
			),
		)
	case !parsedData.GoNext:
		replyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
			),
		)
	}

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		text,
	)
	msg.BaseEdit.ReplyMarkup = &replyMarkup

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("railwaystation.Commander.List: error sending reply message: %v", err)
	}
}
