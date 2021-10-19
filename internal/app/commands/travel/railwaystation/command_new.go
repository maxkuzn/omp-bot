package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Here should be a new command",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("railwaystation.Commander.New: error sending reply message to chat: %v", err)
	}
}
