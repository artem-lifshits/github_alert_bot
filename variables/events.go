package variables

import "github.com/go-playground/webhooks/v6/github"

var GithubEvents = []github.Event{
	github.PushEvent,
	github.PullRequestEvent,
	github.CommitCommentEvent,
	github.PingEvent,
}
