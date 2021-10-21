package railwaystation

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const listLimit = 5

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	stations, err := c.service.List(uint64(0), listLimit)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("An error occured: %s", err))
		return
	}
	if len(stations) == 0 {
		reply(c.bot, inputMessage.Chat.ID, "There are no stations")
		return
	}

	text := "List of stations:\n"
	for _, s := range stations {
		text += fmt.Sprintln(s)
	}

	serializedDataPrev, err := json.Marshal(CallbackListData{
		LastID: 0,
		GoNext: false,
	})
	if err != nil {
		panic(err)
	}
	callbackPathPrev := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "railway_station",
		CallbackName: "list",
		CallbackData: string(serializedDataPrev),
	}

	serializedDataNext, err := json.Marshal(CallbackListData{
		LastID: stations[len(stations)-1].ID,
		GoNext: true,
	})
	if err != nil {
		panic(err)
	}
	callbackPathNext := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "railway_station",
		CallbackName: "list",
		CallbackData: string(serializedDataNext),
	}

	log.Printf("%q = %d", callbackPathPrev.String(), len(callbackPathPrev.String()))

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPathPrev.String()),
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPathNext.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("railwaystation.Commander.List: error sending reply message: %v", err)
	}
}
