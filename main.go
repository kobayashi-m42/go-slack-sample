package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

func main() {
	webookURL := os.Getenv("WEBHOOK_URL")

	arg := strings.Join(os.Args[1:], "")
	params := slack{
		Text:      fmt.Sprintf("%s", arg),
		Username:  "go-slack-sample",
		IconEmoji: ":gopher:",
		IconURL:   "",
		Channel:   "",
	}
	jsonParams, _ := json.Marshal(params)
	resp, _ := http.PostForm(
		webookURL,
		url.Values{"payload": {string(jsonParams)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
