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

	if config.proxyURL != "" {
		dialer, err := proxy.SOCKS5(
			"tcp",
			config.proxyURL,
			&proxy.Auth{User: config.proxyUser, Password: config.proxyPass},
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

	embeddedBot, err := tb.NewBot(tb.Settings{
		Token:  config.tgBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	})

	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	// jira initialization

	tp := jira.BasicAuthTransport{
		Username: config.jiraUser,
		Password: config.jiraPass,
	}
	embeddedJira, err := jira.NewClient(tp.Client(), config.jiraURL)
	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	return &bot{
		embeddedBot,
		embeddedJira,
	}

}
