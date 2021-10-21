package railwaystation

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	usage := "Usage of get command:\n" +
		"/get__travel__railway_station N\n\n" +
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

	station, err := c.service.Describe(stationID)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("An error occured: %s", err))
		return
	}

	reply(c.bot, inputMessage.Chat.ID, fmt.Sprintln(station))
}
