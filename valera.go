package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	PROXY_URL          = "grsst.s5.opennetwork.cc:999"
	PROXY_USER         = "41591017"
	PROXY_PASS         = "5NEISabl"
	TELEGRAM_BOT_TOKEN = "781350316:AAGX5fxcytfyRNlof5SzbJrl42jtHbxFBqI"
)

func main() {

	b := initializeBot()
	b.Raw("deleteWebhook", map[string]string{})

	b.Handle(tb.OnText, b.textMessageHandler)

	b.Start()

}
