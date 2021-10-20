package railwaystation

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	LastID uint64 `json:"last_id"`
}

func (c *Commander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	log.Printf("CallbackList")
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("railwaystation.Commander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	stations, err := c.service.List(parsedData.LastID, listLimit)
	if err != nil {
		reply(c.bot, callback.Message.Chat.ID, fmt.Sprintf("An error occured: %s", err))
		return
	}
	if len(stations) == 0 {
		reply(c.bot, callback.Message.Chat.ID, "There are no more stations")
		return
	}

	text := "List of stations:\n"
	for _, s := range stations {
		text += fmt.Sprintln(s)
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)

	serializedData, err := json.Marshal(CallbackListData{
		LastID: stations[len(stations)-1].ID,
	})
	if err != nil {
		panic(err)
	}

	callbackPath.CallbackData = string(serializedData)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("railwaystation.Commander.List: error sending reply message: %v", err)
	}
}
