package handlers

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-events-go/src/aqua"
	"github.com/BryanKMorrow/aqua-events-go/src/slack"
	"log"
	"net/http"
	"os"
	"strings"
)

// IndexHandler - Home route
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to aqua-events-go"))
}

// SlackHandler handles all slack messages
func SlackHandler(w http.ResponseWriter, r *http.Request) {
	var audit aqua.Audit
	var m slack.Message
	err := json.NewDecoder(r.Body).Decode(&audit)
	if err != nil {
		log.Println("Failed to decode audit event from Aqua")
		log.Println(r.Body)
	}
	w.Header().Set("Content-Type", "application/json")
	webhook, ignore := getEnv()
	m.Webhook = webhook
	m.IgnoreList = ignore
	m.ProcessAudit(audit)
}

// getEnv is an ugly call to get the environment variables that need to be passed
func getEnv() (string, []string) {
	webhook := os.Getenv("SLACK_WEBHOOK")
	ignore := os.Getenv("IGNORE_LIST")
	var splits []string
	if ignore != "" {
		// convert CSV list to slice
		splits = strings.Split(ignore, ",")
	}
	return webhook, splits
}
