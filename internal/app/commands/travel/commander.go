package travel

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/travel/railwaystation"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type TravelCommander struct {
	bot                     *tgbotapi.BotAPI
	railwayStationCommander Commander
}

func NewTravelCommander(
	bot *tgbotapi.BotAPI,
) *TravelCommander {
	return &TravelCommander{
		bot:                     bot,
		railwayStationCommander: railwaystation.NewRailwayStationCommander(bot),
	}
}

func (c *TravelCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "railway_station":
		c.railwayStationCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *TravelCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "railway_station":
		c.railwayStationCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
