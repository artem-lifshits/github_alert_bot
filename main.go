package main

import (
	"fmt"
	"github.com/artem-lifshits/github_alert_bot/pkg/handlers"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {
	token := os.Getenv("TELEGRAM_APITOKEN")
	channel := os.Getenv("TELEGRAM_CHANNEL")
	chatId, err := strconv.ParseInt(channel, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	hook, err := github.New(github.Options.Secret(os.Getenv("GITHUB_SECRET")))
	if err != nil {
		log.Fatal("wrong secret addressed: %w", err)
	}
	newHook := handlers.NewHook(hook, token, chatId)

	handlers.NewHandlers(newHook)

	http.HandleFunc(path, handlers.WebhookCatcher)
	err = http.ListenAndServe(":3535", nil)
	if err != nil {
		fmt.Printf("error occured on server: %s", err.Error())
	}
}
