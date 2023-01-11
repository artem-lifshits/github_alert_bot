package conv

import (
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
)

func ConvertPushPayload(payload interface{}) (msg string) {
	body := payload.(github.PushPayload)
	msg = fmt.Sprintf("New push by %s, created at: %d, commits: %v", body.Pusher.Name, body.Repository.CreatedAt, body.Commits)
	return
}

func ConvertPullRequestPayload(payload interface{}) (msg string) {
	body := payload.(github.PullRequestPayload)
	msg = fmt.Sprintf("%v", body)
	return
}

func ConvertPingPayload(payload interface{}) (msg string) {
	body := payload.(github.PingPayload)
	msg = fmt.Sprintf("%v", body)

	return
}

func ConvertCommitCommentPayload(payload interface{}) (msg string) {
	body := payload.(github.CommitCommentPayload)
	msg = fmt.Sprintf("%v", body)
	return
}
