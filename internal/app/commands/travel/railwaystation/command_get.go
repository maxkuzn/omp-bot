package railwaystation

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RailwayStationCommander) Get(inputMessage *tgbotapi.Message) {
	reply := func(text string) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("RailwayStationCommander.Get: error replying reply message to chat: %v", err)
		}
	}
	usage := "Usage of get command:\n" +
		"/get__travel__railway_station N\n\n" +
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

	station, err := c.service.Describe(stationID)
	if err != nil {
		reply(fmt.Sprintf("An error occured: %s", err))
		return
	}

	reply(fmt.Sprintln(station))
}
