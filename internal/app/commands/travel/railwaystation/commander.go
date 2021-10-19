package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/travel/railwaystation"
)

type RailwayStationCommander struct {
	bot     *tgbotapi.BotAPI
	service railwaystation.RailwayStationService
}

func NewRailwayStationCommander(bot *tgbotapi.BotAPI) *RailwayStationCommander {
	railwayStationService := railwaystation.NewDummyRailwayStationService()
	return &RailwayStationCommander{
		bot:     bot,
		service: railwayStationService,
	}
}

func (c *RailwayStationCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	/*
		case "list":
			c.CallbackList(callback, callbackPath)
	*/
	default:
		log.Printf("RailwayStationCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RailwayStationCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "new":
		c.New(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "edit":
		c.Edit(msg)
	case "delete":
		c.Delete(msg)
	default:
		c.Default(msg)
	}
}
