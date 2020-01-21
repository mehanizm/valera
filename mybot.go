package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/net/proxy"
	tb "gopkg.in/tucnak/telebot.v2"
)

type myBot struct {
	*tb.Bot
}

func initializeBot() *myBot {

	client := http.DefaultClient

	if PROXY_URL != "" {
		dialer, err := proxy.SOCKS5(
			"tcp",
			PROXY_URL,
			&proxy.Auth{User: PROXY_USER, Password: PROXY_PASS},
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
		Token:  TELEGRAM_BOT_TOKEN,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: client,
	})

	if err != nil {
		log.Println("Can't initialize bot", err)
		panic(err)
	}

	return &myBot{
		embededBot,
	}

}
