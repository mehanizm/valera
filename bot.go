package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	jira "github.com/andygrunwald/go-jira"
	"golang.org/x/net/proxy"
	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	*tb.Bot
	*jira.Client
}

func initializeBot(config *config) *bot {

	// telegram initialization

	client := http.DefaultClient

	if config.PROXY_URL != "" {
		dialer, err := proxy.SOCKS5(
			"tcp",
			config.PROXY_URL,
			&proxy.Auth{User: config.PROXY_USER, Password: config.PROXY_PASS},
			proxy.Direct,
		)

		if err != nil {
			log.Panicf("Error in proxy %s", err)
		}

		client = &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.Dial(network, addr)
				},
			},
		}

	}

	embededBot, err := tb.NewBot(tb.Settings{
		Token:  config.TELEGRAM_BOT_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	})

	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	// jira initialization

	tp := jira.BasicAuthTransport{
		Username: config.JIRA_USER,
		Password: config.JIRA_PASS,
	}
	embededJira, err := jira.NewClient(tp.Client(), config.JIRA_URL)
	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	return &bot{
		embededBot,
		embededJira,
	}

}
