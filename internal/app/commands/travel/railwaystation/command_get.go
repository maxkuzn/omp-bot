package railwaystation

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RailwayStationCommander) Get(inputMessage *tgbotapi.Message) {
	send := func(text string) {
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			text,
		)

		_, err := c.bot.Send(msg)
		if err != nil {
			log.Printf("RailwayStationCommander.Get: error sending reply message to chat: %v", err)
		}
	}
	usage := "Usage of get command:\n" +
		"/get__travel__railway_station N\n\n" +
		"N must be a number"

	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	fields := strings.Fields(inputMessage.Text)
	if len(fields) != 2 {
		send(usage)
		return
	}

	stationID, err := strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		send(usage)
		return
	}

	station, err := c.service.Describe(stationID)
	if err != nil {
		send(fmt.Sprintf("An error occured: %s", err))
		return
	}

	send(fmt.Sprintf("Station with id %d\nName: %s\nLocation: %s",
		station.ID, station.Name, station.Location))
}
