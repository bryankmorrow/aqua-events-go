package slack

import (
	"encoding/json"
	"github.com/BryanKMorrow/aqua-events-go/src/aqua"
	"github.com/slack-go/slack"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	AuthorName    = "aqua-events"
	Fallback      = "Aqua Security Audit Events"
	AuthorSubname = "AquaEvents"
	AuthorLink    = "https://github.com/BryanKMorrow/aqua-events-go"
)

type Message struct {
	Attachment  slack.Attachment     `json:"attachment"`
}

func (m *Message) ProcessAudit(audit aqua.Audit) {
	text, err := json.Marshal(&audit)
	if err != nil {
		log.Println(err)
		return
	}
	if audit.Level == "block" {
		m.Attachment.Color = "bad"
	} else if audit.Level == "success" {
		m.Attachment.Color = "good"
	} else {
		m.Attachment.Color = "warn"
	}

	m.Attachment.Fallback = Fallback
	m.Attachment.AuthorName = AuthorName
	m.Attachment.AuthorSubname = AuthorSubname
	m.Attachment.AuthorLink = AuthorLink
	//m.Attachment.AuthorIcon = AuthorIcon
	m.Attachment.Text = string(text)
	m.Attachment.Ts = json.Number(strconv.FormatInt(time.Now().Unix(), 10))
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{m.Attachment},
	}

	err = slack.PostWebhook(os.Getenv("SLACK_WEBHOOK"), &msg)
	if err != nil {
		log.Println("failed posting attachment to Slack API: %w", err)
	}
}


