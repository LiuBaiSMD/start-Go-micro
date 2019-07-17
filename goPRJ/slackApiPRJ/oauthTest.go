package main

//import (
//	"fmt"
//	"github.com/nlopes/slack"
//	"net/http"
//)
//
////import {
////"fmt"
//////"github.com/nlopes/slack"
////}
//func main() {
//	fmt.Println("start")
//
//	api := slack.GetOAuthToken(http.Client, "689344750534.687235193125", "2bf82c709b3432c0d2df99bd2ada81c3", )
//	fmt.Println(rsp, err)
//}

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
)

func getToken(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := []byte(`{"ok": true, "team": {
			"id": "evan",
			"name": "notalar",
          }
		}}`)
	rw.Write(response)
}

func main() {
	http.HandleFunc("/oauth.access", getToken)
	authCode := "***********"
	clientID := "***********"
	clientSecret := "***********"
	redirectURI := ""
	fmt.Println("start")
	oAuthRes, err := slack.GetOAuthResponse(http.DefaultClient, clientID, clientSecret, authCode, redirectURI)
	if err != nil {
		//t.Errorf("Unexpected error: %s", err)
		fmt.Println("error",err)
		return
	}
	fmt.Printf("auth result : %v", oAuthRes)
}