package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxp-689344750534-686762631748-696647166822-01946982c1f49afdd5976302d5a78ee9")
	user, err := api.GetUserInfo("UL6NEJKN0")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}