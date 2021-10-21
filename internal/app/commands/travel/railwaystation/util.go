package railwaystation

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
	"unicode"

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

func splitArgs(data string) []string {
	quoted := false
	args := strings.FieldsFunc(data, func(r rune) bool {
		if r == '"' {
			quoted = !quoted
		}
		return !quoted && unicode.IsSpace(r)
	})
	return args
}

func parseRailwayStationArguments(data string, withID bool) (station travel.RailwayStation, err error) {
	args := splitArgs(data)

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

func parseRailwayStationJSON(data []byte, withID bool) (station travel.RailwayStation, err error) {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	err = d.Decode(&station)
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

func parseRailwayStation(data string, withID bool) (station travel.RailwayStation, err error) {
	data = replaceQuotes(data)
	rawData := []byte(data)
	if json.Valid(rawData) {
		return parseRailwayStationJSON(rawData, withID)
	}
	return parseRailwayStationArguments(data, withID)
}
