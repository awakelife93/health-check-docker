package worker

import (
	"errors"

	"github.com/slack-go/slack"
)

// * Just pass the token over.
var api = slack.New("")

func SendMessage(message string) (string, error) {
	const channelId string = ""

	channelID, timestamp, error := api.PostMessage(
		channelId,
		slack.MsgOptionText(message, false),
	)

	if error != nil {
		return "", errors.New(error.Error())
	}

	return channelID + "/" + timestamp, nil
}
