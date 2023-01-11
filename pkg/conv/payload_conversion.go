package conv

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
)

func ConvertPushPayload(payload interface{}) (msg string) {
	body := payload.(github.PushPayload)
	msg = fmt.Sprintf("New push by '%s'\nCreated at: %s\nCommit: %v", body.Pusher.Name, body.HeadCommit.Timestamp, body.HeadCommit.URL)
	return
}

func ConvertPullRequestPayload(payload interface{}) (msg string) {
	body := payload.(github.PullRequestPayload)
	msg = fmt.Sprintf("New pull request\nCreated at: %s\nUrl: %s", body.PullRequest.CreatedAt, body.PullRequest.URL)
	return
}

func ConvertPingPayload(payload interface{}) (msg string) {
	body := payload.(github.PingPayload)
	msg = fmt.Sprintf("New ping sent at %s", body.Hook.CreatedAt)
	return
}

func ConvertCommitCommentPayload(payload interface{}) (msg string) {
	body := payload.(github.CommitCommentPayload)
	msg = fmt.Sprintf("New commit comment\nBy %s\nCreated at %s\nMessage: %s", body.Comment.User.Login, body.Comment.CreatedAt, body.Comment.Body)
	return
}
