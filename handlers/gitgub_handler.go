package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type githubMessage struct {
	HookId     int        `json:"hook_id"`
	Repository Repository `json:"repository"`
}

type Repository struct {
	Name    string `json:"name"`
	HtmlUrl string `json:"html_url"`
}

func WebhookCatcher(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	log.Println("Incomming Request")
	log.Println("Process JSON")
	decoder := json.NewDecoder(r.Body)
	var t githubMessage

	err := decoder.Decode(&t)
	if err != nil {
		log.Println("Error: ", err)
	}
	//-- OutputEvent Sources
	log.Printf("%v", t)
	headers := r.Header
	log.Println(headers["X-Github-Event"][0])

}
