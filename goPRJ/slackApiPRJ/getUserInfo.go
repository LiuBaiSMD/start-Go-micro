package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("xoxp-689344750534-686762631748-684228523555-7969f76dc4314223ff7fa64965402bd9")
	user, err := api.GetUserInfo("UL6NEJKN0")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}