package main

import (
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *myBot) textMessageHandler(m *tb.Message) {
	log.Printf("Message from %s. Text: %s\n", m.Sender.Username, m.Text)
	b.Send(m.Sender, m.Text)
}
