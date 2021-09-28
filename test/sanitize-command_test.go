package test

import (
	"fmt"
	"github.com/endriu00/DiscordBuilderBot/service"
	"github.com/sirupsen/logrus"
	"testing"
)

const correctCommand = "!golang"
const correctCommandAnswer = "golang"
const wrongCommand0 = "golang"
const wrongCommand1 = "/golang"

func TestSanitizeCommand(t *testing.T) {
	fakeConfig := service.Config{
		Log: logrus.NewEntry(logrus.StandardLogger()),
	}
	bot := service.New(fakeConfig)

	// Define tests
	var tests = []struct {
		command string
		want    string
	}{
		{correctCommand, correctCommandAnswer},
		{wrongCommand0, ""},
		{wrongCommand1, ""},
	}

	// Begin testing
	for _, test := range tests {
		testName := fmt.Sprintf("%s,%s", test.command, test.want)
		t.Run(testName, func(t *testing.T) {
			answer, _ := bot.SanitizeCommand(test.command)
			if answer != test.want {
				t.Errorf("Got %s, wanted %s", answer, test.want)
			}
		})
	}
}
