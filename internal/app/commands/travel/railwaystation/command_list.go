package railwaystation

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	reply := func(text string) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("railwaystation.Commander.List: error replying reply message to chat: %v", err)
		}
	}

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	stations, err := c.service.List(uint64(0), 100)
	if err != nil {
		reply(fmt.Sprintf("An error occured: %s", err))
		return
	}

	text := "List of stations:\n"
	for _, s := range stations {
		text += fmt.Sprintln(s)
	}
	reply(text)
}
