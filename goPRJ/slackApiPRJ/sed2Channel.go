package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxp-689344750534-686762631748-696647166822-01946982c1f49afdd5976302d5a78ee9", slack.OptionDebug(true))
	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
		ImageURL:"https://emoji.slack-edge.com/TL9A4N2FQ/cxk/9ea39b9eb3c84d41.png",
		// Uncomment the following part to send a field too
		//	Fields: []slack.AttachmentField{
		//		slack.AttachmentField{
		//			Title: "a",
		//			Value: "no",
		//		},
		//	},
	}

	channelID, timestamp, err := api.PostMessage("CL76N8WSC", slack.MsgOptionText("Some text", false), slack.MsgOptionAttachments(attachment))
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}