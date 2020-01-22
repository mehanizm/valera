package main

import (
	"strings"

	tb "gopkg.in/tucnak/telebot.v2"
)

type handlerFunc func(m *tb.Message)

func (b *bot) textMessageHandler() func(m *tb.Message) {
	return func(m *tb.Message) {

		taskKeys := parseIssueKeysFromMsg(m.Text)

		if len(taskKeys) == 0 {
			return
		}

		taskDescriptions := make([]string, len(taskKeys))
		for _, taskKey := range taskKeys {
			taskDescriptions = append(taskDescriptions, b.getIssueDescFromJira(taskKey))
		}
		messageText := strings.Join(taskDescriptions, "\n")
		b.Send(m.Chat, messageText)

	}
}
