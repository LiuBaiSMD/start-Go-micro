package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	var (
		apiToken string
		debug    bool
	)
//xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r
	flag.StringVar(&apiToken, "token", "xoxp-689344750534-686762631748-696647166822-01946982c1f49afdd5976302d5a78ee9", "test")
	flag.BoolVar(&debug, "debug", false, "Show JSON output")
	flag.Parse()

	api := slack.New(apiToken, slack.OptionDebug(debug))

	var (
		postAsUserName  string
		postAsUserID    string
		postToUserName  string
		postToUserID    string
		postToChannelID string
	)

	// Find the user to post as.
	authTest, err := api.AuthTest()
	if err != nil {
		fmt.Printf("Error getting channels: %s\n", err)
		return
	}

	// Post as the authenticated user.
	postAsUserName = authTest.User
	postAsUserID = authTest.UserID

	// Posting to DM with self causes a conversation with slackApiPRJ.
	postToUserName = authTest.User
	postToUserID = authTest.UserID

	// Find the channel.
	_, _, chanID, err := api.OpenIMChannel(postToUserID)
	if err != nil {
		fmt.Printf("Error opening IM: %s\n", err)
		return
	}
	postToChannelID = chanID

	fmt.Printf("Posting as %s (%s) in DM with %s (%s), channel %s\n", postAsUserName, postAsUserID, postToUserName, postToUserID, postToChannelID)

	// Post a message.
	channelID, timestamp, err := api.PostMessage(postToChannelID, slack.MsgOptionText("Is this any good?", false))
	if err != nil {
		fmt.Printf("Error posting message: %s\n", err)
		return
	}

	// Grab a reference to the message.
	msgRef := slack.NewRefToMessage(channelID, timestamp)

	if err = api.AddReaction("cxk", msgRef); err != nil {
		fmt.Printf("Error adding reaction: %s\n", err)
		return
	}
}
