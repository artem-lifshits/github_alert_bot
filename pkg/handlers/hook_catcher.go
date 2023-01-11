package handlers

import (
	"fmt"
	"github.com/artem-lifshits/github_alert_bot/pkg/conv"
	vars "github.com/artem-lifshits/github_alert_bot/variables"
	"github.com/go-playground/webhooks/v6/github"
	"log"
	"net/http"
)

var Repo *Hook

type Hook struct {
	Webhook *github.Webhook
	Token   string
	Channel int64
}

func NewHook(a *github.Webhook, b string, c int64) *Hook {
	return &Hook{
		Webhook: a,
		Token:   b,
		Channel: c,
	}
}

func NewHandlers(r *Hook) {
	Repo = r
}

func WebhookCatcher(_ http.ResponseWriter, r *http.Request) {
	fmt.Println("Request incoming")
	payload, err := Repo.Webhook.Parse(r, vars.GithubEvents...)
	if err != nil {
		if err == github.ErrEventNotFound {
			fmt.Println("Incoming hook is not part of github events listed above")
			return
		}
	}
	var message string

	switch payload := payload.(type) {

	case github.PushPayload:
		message = conv.ConvertPushPayload(payload)

	case github.PullRequestPayload:
		message = conv.ConvertPullRequestPayload(payload)

	case github.PingPayload:
		message = conv.ConvertPingPayload(payload)

	case github.CommitCommentPayload:
		message = conv.ConvertCommitCommentPayload(payload)
	}

	chanMessage := Message{
		ChatID: Repo.Channel,
		Text:   message,
	}

	err = SendMessage(FormUrl(), &chanMessage)
	if err != nil {
		log.Fatal(err)
	}
}
