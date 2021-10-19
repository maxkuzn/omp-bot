package railwaystation

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/travel/railwaystation"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service railwaystation.Service
}

func NewCommander(bot *tgbotapi.BotAPI) *Commander {
	return &Commander{
		bot:     bot,
		service: railwaystation.NewDummyService(),
	}
}

func (c *Commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	/*
		case "list":
			c.CallbackList(callback, callbackPath)
	*/
	default:
		log.Printf("railwaystation.Commander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *Commander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
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
