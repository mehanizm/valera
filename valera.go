package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	configFilePath       = flag.String("config", "configs/config.yaml", "Configuration file path")
	allowedChatsFilePath = flag.String("allowed", "configs/whitelist.txt", "White list of the chats")
)

func init() {
	// Log as TEXT
	log.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		HideKeys:        true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

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

	b, err := initializeBot(config)
	if err != nil {
		os.Exit(2)
	}
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
