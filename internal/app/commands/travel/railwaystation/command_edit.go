package railwaystation

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) Edit(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	station, err := parseRailwayStation(inputMessage.CommandArguments(), true)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Error parsing arguments: %v", err))
		return
	}

	err = c.service.Update(station.ID, station)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Error updating station: %v", err))
	}

	reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Succsessfully updated station with id %d", station.ID))
}
