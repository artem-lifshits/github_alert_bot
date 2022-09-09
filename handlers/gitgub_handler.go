package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type githubMessage struct {
	Action string `json:"action"`
}

func WebhookCatcher(w http.ResponseWriter, r *http.Request) {
	log.Println("Incomming Request")
	log.Println("Process JSON")
	decoder := json.NewDecoder(r.Body)
	var t githubMessage

	err := decoder.Decode(&t)
	if err != nil {
		log.Println("Error: ", err)
	}
	//-- OutputEvent Source
	log.Println("Action Name ", t.Action)
}
