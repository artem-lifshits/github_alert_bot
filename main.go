package main

import (
	"fmt"
	"github.com/artem-lifshits/github_alert_bot/pkg/handlers"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/webhooks/v6/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, err := github.New(github.Options.Secret(os.Getenv("GITHUB_SECRET")))
	if err != nil {
		log.Fatal("wrong secret addressed: %w", err)
	}
	newHook := handlers.NewHook(hook)
	handlers.NewHandlers(newHook)

	http.HandleFunc(path, handlers.WebhookCatcher)
	err = http.ListenAndServe(":3535", nil)
	if err != nil {
		fmt.Printf("error occured on server: %s", err.Error())
	}
}
