package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"
	"unicode/utf8"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/proxy"
	tb "gopkg.in/tucnak/telebot.v2"
)

type bot struct {
	*tb.Bot
	*jira.Client
}

// initializeBot
// initialize bot structure by the config (telegram and jira)
// configure proxy if it exists in the config
func initializeBot(config *config) (*bot, error) {

	if config == nil {
		emptyConfigError := errors.New("There is empty config struct")
		log.WithField("component", "initialize bot").Error(emptyConfigError)
		return nil, emptyConfigError
	}

	if utf8.RuneCountInString(config.TgBotToken) != 45 {
		notValidTokenError := errors.New("Telegram bot token is not valid")
		log.WithField("component", "initialize bot").Error(notValidTokenError)
		return nil, notValidTokenError
	}

	// telegram initialization

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	if config.ProxyURL != "" {
		dialer, err := proxy.SOCKS5(
			"tcp",
			config.ProxyURL,
			&proxy.Auth{User: config.ProxyUser, Password: config.ProxyPass},
			proxy.Direct,
		)

		if err != nil {
			log.WithField("component", "initialize proxy").Error(err)
			return nil, err
		}

		client = &http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return dialer.Dial(network, addr)
				},
			},
		}

		log.WithField("component", "initialize proxy").
			Infof("connecting to telegram with proxy %v", config.ProxyURL)

	}

	embeddedBot, err := tb.NewBot(tb.Settings{
		Token:  config.TgBotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	})

	if err != nil {
		log.WithField("component", "initialize bot").Error(err)
		return nil, err
	}

	embeddedBot.Raw("deleteWebhook", map[string]string{})

	// jira initialization

	tp := jira.BasicAuthTransport{
		Username: config.JiraUser,
		Password: config.JiraPass,
	}
	embeddedJira, err := jira.NewClient(tp.Client(), config.JiraURL)
	if err != nil {
		log.WithField("component", "initialize jira").Error(err)
		return nil, err
	}
	_, _, err = embeddedJira.User.GetSelf()
	if err != nil {
		log.WithField("component", "initialize jira").Error(err)
		return nil, err
	}

	return &bot{
		embeddedBot,
		embeddedJira,
	}, nil

}
