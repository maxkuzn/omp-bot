package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	outputText := ""
	outputText += "/help__travel__railway_station - print list of commands\n"
	outputText += "/new__travel__railway_station - create a new info about railway station\n"
	outputText += "/get__travel__railway_station - get an info about railway station\n"
	outputText += "/list__travel__railway_station - get a list of all railway stations\n"
	outputText += "/edit__travel__railway_station - edit an info about railway station\n"
	outputText += "/delete__travel__railway_station - delete a railway station\n"

	reply(c.bot, inputMessage.Chat.ID, outputText)
}
