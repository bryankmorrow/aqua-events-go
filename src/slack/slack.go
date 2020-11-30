package slack

import (
	"encoding/json"
	"fmt"
	"github.com/BryanKMorrow/aqua-events-go/src/aqua"
	"github.com/slack-go/slack"
	"log"
	"strconv"
	"time"
)

const (
	// AuthorName is the message identifier
	AuthorName = "aqua-events"
	// Fallback is the backup for AuthorName
	Fallback = "Aqua Security Audit Events"
	// AuthorSubname follows the AuthorName in the header
	AuthorSubname = "AquaEvents"
	// AuthorLink points to the github repo for this application
	AuthorLink = "https://github.com/BryanKMorrow/aqua-events-go"
	// AuthorIcon points to the Aqua favicon
	AuthorIcon = "https://www.aquasec.com/wp-content/themes/aqua3/favicon.ico"
)

// Message is the slack struct
type Message struct {
	Attachment slack.Attachment `json:"attachment"`
	Webhook    string           `json:"webhook"`
	IgnoreList []string         `json:"ignore_list"`
}

// ProcessAudit receives the post data and sends to slack
func (m *Message) ProcessAudit(audit aqua.Audit) {
	// format the message
	ignore := false
	msg := m.Format(audit)

	if audit.Result == 2 { // BLOCK
		contains := sliceContains(m.IgnoreList, "block")
		if contains {
			log.Println("ignoring block events")
			ignore = true
		}
	} else if audit.Result == 1 { // SUCCESS
		contains := sliceContains(m.IgnoreList, "success")
		if contains {
			log.Println("ignoring success events")
			ignore = true
		}
	} else if audit.Result == 3 { // DETECT
		contains := sliceContains(m.IgnoreList, "detect")
		if contains {
			log.Println("ignoring detect events")
			ignore = true
		}
	} else if audit.Result == 4 {
		contains := sliceContains(m.IgnoreList, "alert")
		if contains {
			log.Println("ignoring critical events")
			ignore = true
		}
	}

	if !ignore {
		err := slack.PostWebhook(m.Webhook, &msg)
		if err != nil {
			log.Println("failed posting attachment to Slack API: %w", err)
		}
	}
}

func (m *Message) Format(audit aqua.Audit) slack.WebhookMessage {
	var text string
	var err error
	var data []byte
	// base attachment settings
	m.Attachment.Fallback = Fallback
	m.Attachment.AuthorName = AuthorName
	m.Attachment.AuthorSubname = AuthorSubname
	m.Attachment.AuthorLink = AuthorLink
	m.Attachment.AuthorIcon = AuthorIcon
	// format based on message level
	if audit.Result == 1 {
		m.Attachment.Color = "good"
		if audit.Type == "Administration" {
			text = fmt.Sprintf("Type: %s\nAction: %s\nPerformed On: %s\nPerformed By: %s\nAqua Response: %s\nTimestamp: %s\n",
				audit.Type, audit.Action, fmt.Sprintf("%s %s", audit.Category, audit.Adjective), audit.User, "Success", time.Unix(int64(audit.Time), 0).Format(time.RFC822Z))
			m.Attachment.AuthorSubname = fmt.Sprintf("User %s performed %s on %s", audit.User, audit.Action, fmt.Sprintf("%s %s", audit.Category, audit.Adjective))
		} else if audit.Type == "CVE" || audit.Category == "CVE" {
			text = fmt.Sprintf("Image: %s\nImage Hash: %s\nRegistry: %s\nImage added by user: %s\nImage scan start time: %s\nImage scan end time: %s\nAqua Response: %s\nTimestamp: %s\n",
				audit.Image, audit.Imagehash, audit.Registry, audit.User, time.Unix(int64(audit.StartTime), 0).Format(time.RFC822Z), time.Unix(int64(audit.Time), 0).Format(time.RFC822Z),
				"Success", time.Unix(int64(audit.Time), 0).Format(time.RFC822Z))
			m.Attachment.AuthorSubname = fmt.Sprintf("Scan of image %s from registry %s revealed critical:%d high:%d medium:%d low:%d", audit.Image, audit.Registry, audit.Critical, audit.High,
				audit.Medium, audit.Low)
		} else if audit.Type == "Docker" || audit.Category == "container" || audit.Category == "image" {
			text = fmt.Sprintf("Host: %s\nHost IP: %s\nImage Name: %s\nContainer Name: %s\nAction: %s\nKubernetes Cluster: %s\nVM Location: %s\nAqua Response: %s\nAqua Policy: %s\nDetails: %s\n"+
				"Enforcer Group: %s\nTime Stamp: %s\n", audit.Host, audit.Hostip, audit.Image, audit.Container, audit.Action, audit.K8SCluster, audit.VMLocation, "Success", audit.Rule, audit.Reason,
				audit.Hostgroup, time.Unix(int64(audit.Time), 0).Format(time.RFC822Z))
			m.Attachment.AuthorSubname = fmt.Sprintf("User ran container %s on host %s", audit.Action, audit.Host)
		} else {
			data, err = json.Marshal(&audit)
			if err != nil {
				log.Println(err)
			}
			text = string(data)
		}
	} else if audit.Result == 3 {
		m.Attachment.Color = "warning"
		data, err = json.Marshal(&audit)
		if err != nil {
			log.Println(err)
		}
		text = string(data)
	} else if audit.Result == 2 {
		m.Attachment.Color = "danger"
		data, err = json.Marshal(&audit)
		if err != nil {
			log.Println(err)
		}
		text = string(data)
	} else if audit.Result == 4 {
		m.Attachment.Color = "danger"
		data, err = json.Marshal(&audit)
		if err != nil {
			log.Println(err)
		}
		text = string(data)
	}

	m.Attachment.Text = text
	m.Attachment.Ts = json.Number(strconv.FormatInt(time.Now().Unix(), 10))
	msg := slack.WebhookMessage{
		Attachments: []slack.Attachment{m.Attachment},
	}
	return msg
}

// sliceContains checks for a string in a slice
func sliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
