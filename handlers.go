package main

import (
	"log"
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *bot) textMessageHandler(m *tb.Message) {
	log.Printf("Message from %s. Text: %s\n", m.Sender.Username, m.Text)

	taskKeys := parseIssueKeysFromMsg(m.Text)

	if len(taskKeys) == 0 {
		b.Send(m.Sender, "There are no task keys in the message")
		return
	}

	taskDescriptions := make([]string, len(taskKeys))
	for _, taskKey := range taskKeys {
		taskDescriptions = append(taskDescriptions, b.getIssueDescFromJira(taskKey))
	}
	messageText := strings.Join(taskDescriptions, "\n")
	b.Send(m.Sender, messageText)

}
