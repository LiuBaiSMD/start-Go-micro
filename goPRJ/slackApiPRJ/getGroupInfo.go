package main

import (
	"fmt"
	"github.com/nlopes/slack"
)

func main() {
//api := slack.New("xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r")
// If you set debugging, it will log all requests to the console
// Useful when encountering issues
api := slack.New("xoxp-689344750534-686762631748-696647166822-01946982c1f49afdd5976302d5a78ee9", slack.OptionDebug(true))
groups, err := api.GetGroups(false)
if err != nil {
fmt.Printf("%s\n", err)
return
}
fmt.Println("start")
for _, group := range groups {
fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
}
fmt.Println("over")
}