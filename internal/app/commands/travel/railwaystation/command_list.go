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

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, text)

	serializedData, err := json.Marshal(CallbackListData{
		LastID: stations[len(stations)-1].ID,
	})
	if err != nil {
		panic(err)
	}

	callbackPath := path.CallbackPath{
		Domain:       "travel",
		Subdomain:    "railway_station",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

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
