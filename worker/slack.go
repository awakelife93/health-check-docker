package worker

import (
	"errors"

	"github.com/slack-go/slack"
)

// * slack channel id
var channelId string = ""

// * slack sdk object
var api = slack.New("")

func SendMessage(message string) (string, error) {

	channelID, timestamp, error := api.PostMessage(
		channelId,
		slack.MsgOptionText(message, false),
	)

	if error != nil {
		return "", errors.New(error.Error())
	}

	return channelID + "/" + timestamp, nil
}
