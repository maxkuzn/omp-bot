package railwaystation

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	reply(c.bot, inputMessage.Chat.ID,
		fmt.Sprintf("You wrote %q, but I do not understand it :(\nPlease write /help__travel__railway_station to get usage info.", inputMessage.Text))
}
