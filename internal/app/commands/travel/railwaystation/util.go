package railwaystation

import (
	"encoding/csv"
	"encoding/json"
	"errors"
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
	r := csv.NewReader(strings.NewReader(text))
	r.Comma = ' '
	args, err := r.Read()
	if err != nil {
		return nil, err
	}
	return args, nil
}

func parseRailwayStationArguments(text string, withID bool) (station travel.RailwayStation, err error) {
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

func parseRailwayStationJSON(text string, withID bool) (station travel.RailwayStation, err error) {
	d := json.NewDecoder(strings.NewReader(text))
	d.DisallowUnknownFields()
	err = json.Unmarshal([]byte(text), &station)
	if err != nil {
		return
	}
	if len(station.Name) == 0 {
		err = errors.New("you should specify non-zero name")
		return
	}
	if len(station.Location) == 0 {
		err = errors.New("you should specify non-zero location")
		return
	}
	if withID && station.ID == 0 {
		err = errors.New("you should specify valid id")
		return
	}
	return
}

func parseRailwayStation(text string, withID bool) (station travel.RailwayStation, err error) {
	text = replaceQuotes(text)
	spaceIdx := strings.IndexByte(text, ' ')
	if spaceIdx == -1 {
		err = errors.New("invalid format")
		return
	}
	if json.Valid([]byte(text[spaceIdx+1:])) {
		return parseRailwayStationJSON(text[spaceIdx+1:], withID)
	}
	return parseRailwayStationArguments(text, withID)
}
