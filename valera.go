package main

import (
	"flag"
	"log"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	configFilePath = flag.String("config", "config.yaml", "Путь до конфигурационного файла")
)

func main() {

	flag.Parse()
	config, err := parseConfigFromFile(*configFilePath)
	if err != nil {
		return
	}

	b := initializeBot(config)
	b.Raw("deleteWebhook", map[string]string{})

	b.Handle(tb.OnText, b.textMessageHandler)

	log.Println("Start telegram bot")
	b.Start()

}
