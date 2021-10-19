package railwaystation

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	reply := func(text string) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("railwaystation.Commander.Delete: error replying reply message to chat: %v", err)
		}
	}

	usage := "Usage of delete command:\n" +
		"/delete__travel__railway_station N\n\n" +
		"N - station id (must be a number)"

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	fields := strings.Fields(inputMessage.Text)
	if len(fields) != 2 {
		reply(usage)
		return
	}

	stationID, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		reply(usage)
		return
	}

	ok, err := c.service.Remove(stationID)
	if err != nil {
		reply(fmt.Sprintf("An error occured: %s", err))
		return
	}
	if !ok {
		reply(fmt.Sprintf("Cannot delete station with id %d", stationID))
		return
	}

	reply(fmt.Sprintf("Station with id %d was successfully deleted", stationID))
}
