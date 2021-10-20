package railwaystation

import (
	"encoding/csv"
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

func replaceQuotes(text string) string {
	text = strings.ReplaceAll(text, "“", "\"")
	text = strings.ReplaceAll(text, "”", "\"")
	text = strings.ReplaceAll(text, "‘", "'")
	text = strings.ReplaceAll(text, "’", "'")
	return text
}

func splitArgs(text string) ([]string, error) {
	text = replaceQuotes(text)

	r := csv.NewReader(strings.NewReader(text))
	r.Comma = ' '
	args, err := r.Read()
	if err != nil {
		return nil, err
	}
	return args, nil
}

func parseRailwayStation(text string, withID bool) (station travel.RailwayStation, err error) {
	args, err := splitArgs(text)
	if err != nil {
		return
	}

	fs := flag.NewFlagSet("parse", flag.ContinueOnError)
	builder := strings.Builder{}
	fs.SetOutput(&builder)

	fs.StringVar(&station.Name, "Name", "", "Name of station")
	fs.StringVar(&station.Location, "Location", "", "Location of station")
	if withID {
		fs.Uint64Var(&station.ID, "ID", 0, "Station ID")
	}

	err = fs.Parse(args[1:])
	if err != nil {
		return
	}
	log.Println(station)

	if !((withID && fs.NFlag() == 3) || (!withID && fs.NFlag() == 2)) {
		fs.PrintDefaults()
		err = fmt.Errorf("\nUsage:\n%s", builder.String())
		return
	}
	return
}
