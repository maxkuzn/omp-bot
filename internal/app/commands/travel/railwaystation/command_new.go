package railwaystation

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *Commander) New(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

	fields := strings.Fields(inputMessage.Text)
	station, err := parseRailwayStation(fields[1:], false)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Error parsing arguments: %v", err))
		return
	}

	id, err := c.service.Create(station)
	if err != nil {
		reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Error creating new station: %v", err))
	}

	reply(c.bot, inputMessage.Chat.ID, fmt.Sprintf("Succsessfully created station with id %d", id))
}
