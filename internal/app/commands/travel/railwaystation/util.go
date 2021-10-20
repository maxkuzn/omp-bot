package railwaystation

import (
	"flag"
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/travel"
)

func reply(bot *tgbotapi.BotAPI, chatID int64, text string) {
	msg := tgbotapi.NewMessage(
		chatID,
		text,
	)

	_, err := bot.Send(msg)
	if err != nil {
		log.Printf("error replying reply message to chat: %v", err)
	}
}

func parseRailwayStation(args []string, withID bool) (station travel.RailwayStation, err error) {
	fs := flag.NewFlagSet("parse", flag.ContinueOnError)
	builder := strings.Builder{}
	fs.SetOutput(&builder)

	fs.StringVar(&station.Name, "Name", "", "Name of station")
	fs.StringVar(&station.Location, "Location", "", "Location of station")
	if withID {
		fs.Uint64Var(&station.ID, "ID", 0, "Station ID")
	}

	err = fs.Parse(args)
	if err != nil {
		return
	}

	if !((withID && fs.NFlag() == 3) || (!withID && fs.NFlag() == 2)) {
		fs.PrintDefaults()
		err = fmt.Errorf("\nUsage:\n%s", builder.String())
		return
	}
	return
}
