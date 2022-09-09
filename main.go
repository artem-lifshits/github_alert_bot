package main

import (
	bot "github.com/artem-lifshits/github_alert_bot/bot"
	handlers "github.com/artem-lifshits/github_alert_bot/handlers"
	"log"
	"net/http"
)

func main() {
	bot.InitializeBot()

	http.HandleFunc("/", handlers.WebhookCatcher)

	err = http.ListenAndServe(":3535", nil)
	if err != nil {
		log.Fatal(err)
	}
}
