package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	configFilePath       = flag.String("config", "configs/config.yaml", "Configuration file path")
	allowedChatsFilePath = flag.String("allowed", "configs/whitelist.txt", "White list of the chats")
)

func main() {

	flag.Parse()
	config, err := parseConfigFromFile(*configFilePath)
	if err != nil {
		return
	}

	rand.Seed(time.Now().UnixNano())
	auth := authData{
		secret:       randSeq(30),
		allowedChats: make(map[string]bool, 0),
		filePath:     *allowedChatsFilePath,
	}
	err = auth.readAuthFromFile()
	if err != nil {
		log.Println("ERROR:", err)
	}
	log.Println("To add chat to the white list, please, send the string to the Bot:", auth.secret)

	b := initializeBot(config)
	b.Raw("deleteWebhook", map[string]string{})

	b.Handle(
		tb.OnText,
		logTextMessageMiddleware(
			auth.checkAuthMiddleware(
				b.textMessageHandler(),
			),
		),
	)

	log.Println("Start telegram bot")
	b.Start()

}
