package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	configFilePath       = flag.String("config", "configs/config_test.yaml", "Путь до конфигурационного файла")
	allowedChatsFilePath = flag.String("allowed", "configs/whitelist.txt", "Путь до разрешенных чатов")
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
	log.Println("To add chat to white list, please, send the string to Bot", auth.secret)

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
