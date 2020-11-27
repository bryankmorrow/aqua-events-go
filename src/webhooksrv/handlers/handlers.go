package handlers

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-events-go/src/aqua"
	"github.com/BryanKMorrow/aqua-events-go/src/slack"
	"log"
	"net/http"
)

// Index - Home route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to aqua-events-go"))
}


func SlackHandler(w http.ResponseWriter, r *http.Request) {
	var audit aqua.Audit
	var m slack.Message
	err := json.NewDecoder(r.Body).Decode(&audit)
	if err != nil {
		log.Println("Failed to decode audit event from Aqua")
	}
	w.Header().Set("Content-Type", "application/json")
	m.ProcessAudit(audit)
}
