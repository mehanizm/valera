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

	if config.ProxyURL != "" {
		dialer, err := proxy.SOCKS5(
			"tcp",
			config.ProxyURL,
			&proxy.Auth{User: config.ProxyUser, Password: config.ProxyPass},
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
		Token:  config.TgBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	})

	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	// jira initialization

	tp := jira.BasicAuthTransport{
		Username: config.JiraUser,
		Password: config.JiraPass,
	}
	embeddedJira, err := jira.NewClient(tp.Client(), config.JiraURL)
	if err != nil {
		log.Println("Can't initialize bot", err)
		return nil
	}

	return &bot{
		embeddedBot,
		embeddedJira,
	}

}
