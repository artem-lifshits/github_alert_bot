package handlers

import (
	"fmt"
	"github.com/artem-lifshits/github_alert_bot/pkg/conv"
	vars "github.com/artem-lifshits/github_alert_bot/variables"
	"github.com/go-playground/webhooks/v6/github"
	"net/http"
)

var Repo *Hook

type Hook struct {
	Webhook *github.Webhook
}

func NewHook(a *github.Webhook) *Hook {
	return &Hook{
		Webhook: a,
	}
}

func NewHandlers(r *Hook) {
	Repo = r
}

func WebhookCatcher(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hook incoming")
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
	fmt.Println(message)
}
