package railwaystation

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	stations, err := c.service.List(uint64(0), 100)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("An error occured: %s", err))
		return
	}

	text := "List of stations:\n"
	for _, s := range stations {
		text += fmt.Sprintln(s)
	}
	reply(c.bot, inputMessage.Chat.ID, text)
}
