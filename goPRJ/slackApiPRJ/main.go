package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

//func main() {
////api := slack.New("xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r")
//// If you set debugging, it will log all requests to the console
//// Useful when encountering issues
//api := slack.New("xoxp-689344750534-686762631748-694453601456-bef40129bd0f3a430b14dcfaa1339d41", slack.OptionDebug(true))
//groups, err := api.GetGroups(false)
//if err != nil {
//fmt.Printf("%s\n", err)
//return
//}
//fmt.Println("start")
//for _, group := range groups {
//fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
//}
//fmt.Println("over")
//}


//func main() {
//	api := slack.New("xoxp-689344750534-686762631748-694453601456-bef40129bd0f3a430b14dcfaa1339d41")
//	user, err := api.GetUserInfo("UL6NEJKN0")
//	if err != nil {
//		fmt.Printf("%s\n", err)
//		return
//	}
//	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
//}

//func TestGetUserInfo(t *testing.T) {
//	http.HandleFunc("/users.info", getUserInfo)
//	expectedUser := getTestUser()
//
//	once.Do(startServer)
//	api := New("xoxp-689344750534-686762631748-694453601456-bef40129bd0f3a430b14dcfaa1339d41", OptionAPIURL("http://"+serverAddr+"/"))
//
//	user, err := api.GetUserInfo("UXXXXXXXX")
//	if err != nil {
//		t.Errorf("Unexpected error: %s", err)
//		return
//	}
//	if !reflect.DeepEqual(expectedUser, *user) {
//		t.Fatal(ErrIncorrectResponse)
//	}
//}
//func main(){
//	TestGetUserInfo()
//}


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


//func main() {
//	var (
//		apiToken string
//		debug    bool
//	)
////xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r
//	flag.StringVar(&apiToken, "token", "xoxp-689344750534-686762631748-696647166822-01946982c1f49afdd5976302d5a78ee9", "test")
//	flag.BoolVar(&debug, "debug", false, "Show JSON output")
//	flag.Parse()
//
//	api := slack.New(apiToken, slack.OptionDebug(debug))
//
//	var (
//		postAsUserName  string
//		postAsUserID    string
//		postToUserName  string
//		postToUserID    string
//		postToChannelID string
//	)
//
//	// Find the user to post as.
//	authTest, err := api.AuthTest()
//	if err != nil {
//		fmt.Printf("Error getting channels: %s\n", err)
//		return
//	}
//
//	// Post as the authenticated user.
//	postAsUserName = authTest.User
//	postAsUserID = authTest.UserID
//
//	// Posting to DM with self causes a conversation with slackApiPRJ.
//	postToUserName = authTest.User
//	postToUserID = authTest.UserID
//
//	// Find the channel.
//	_, _, chanID, err := api.OpenIMChannel(postToUserID)
//	if err != nil {
//		fmt.Printf("Error opening IM: %s\n", err)
//		return
//	}
//	postToChannelID = chanID
//
//	fmt.Printf("Posting as %s (%s) in DM with %s (%s), channel %s\n", postAsUserName, postAsUserID, postToUserName, postToUserID, postToChannelID)
//
//	// Post a message.
//	channelID, timestamp, err := api.PostMessage(postToChannelID, slack.MsgOptionText("Is this any good?", false))
//	if err != nil {
//		fmt.Printf("Error posting message: %s\n", err)
//		return
//	}
//
//	// Grab a reference to the message.
//	msgRef := slack.NewRefToMessage(channelID, timestamp)
//
//	if err = api.AddReaction("cxk", msgRef); err != nil {
//		fmt.Printf("Error adding reaction: %s\n", err)
//		return
//	}
//}
