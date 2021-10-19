package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RailwayStationCommander) Delete(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		"Here should be a delete command",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("RailwayStationCommander.Delete: error sending reply message to chat: %v", err)
	}
}
