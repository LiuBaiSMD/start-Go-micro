// xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r
package main

import (
	"fmt"
	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"
	"os"

	// Import all the commands you wish to use
)

func main() {
	//slack.Run("xoxb-689344750534-682886786835-pnr0cMHputWkcUlc0TOqyW1r")
	fmt.Println("----->start")
	token := os.Getenv("SLACK_TOKEN")
	fmt.Println(token)
	slack.Run(token)
}