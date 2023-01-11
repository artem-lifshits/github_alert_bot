package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Message struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func FormUrl() string {
	str := []string{"https://api.telegram.org/bot", Repo.Token, "/sendMessage"}
	url := strings.Join(str, "")
	return url
}

func SendMessage(url string, message *Message) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			log.Println("failed to close response body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status %s", resp.Status)
	}

	return nil
}
