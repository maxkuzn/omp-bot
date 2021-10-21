package railwaystation

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Delete(inputMessage *tgbotapi.Message) {
	usage := "Usage of delete command:\n" +
		"/delete__travel__railway_station N\n\n" +
		"N - station id (must be a number)"

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	fields := strings.Fields(inputMessage.Text)
	if len(fields) != 2 {
		reply(c.bot, inputMessage.Chat.ID, usage)
		return
	}

	stationID, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, usage)
		return
	}

	ok, err := c.service.Remove(stationID)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("An error occured: %s", err))
		return
	}
	if !ok {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Cannot delete station with id %d", stationID))
		return
	}

	reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Station with id %d was successfully deleted", stationID))
}
