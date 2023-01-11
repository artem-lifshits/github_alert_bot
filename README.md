## Github alert bot

This is a simple github alert bot which sends github updates to telegram channel.

4 alert actions are supported at the moment.

Webhook `Payload URL` has a fixed structure: `http://example.com:3535/webhooks`.

### Required environment variables

```sh
$ `export GITHUB_SECRET="xxxxx"`
$ `export TELEGRAM_APITOKEN="xxxxx"`
$ `export TELEGRAM_CHANNEL="xxxxx"`
```
