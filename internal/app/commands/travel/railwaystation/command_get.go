package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RailwayStationCommander) Get(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Here should be a get command",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("RailwayStationCommander.Get: error sending reply message to chat: %v", err)
	}
}
