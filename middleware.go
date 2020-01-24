package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

type authData struct {
	secret       string
	allowedChats map[string]bool
	filePath     string
}

// saveAuthToFile save data about
// allowed chat that collected in memory
// to file
func (a *authData) saveAuthToFile() error {

	f, err := os.Create(a.filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for chatID := range a.allowedChats {
		fmt.Fprintln(w, chatID)
	}
	w.Flush()

	log.WithField("component", "auth saver").Infof("Data saved to file %v", a.filePath)

	return nil
}

// readAuthFromFile read file with
// the white list chats if it exists
func (a *authData) readAuthFromFile() error {

	f, err := os.Open(a.filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		a.allowedChats[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	log.WithField("component", "auth reader").
		Infof("Successfully read %v lines from auth file", len(a.allowedChats))

	return nil
}

func (a *authData) checkAuthMiddleware(next func(m *tb.Message)) func(m *tb.Message) {
	return func(m *tb.Message) {

		switch {
		case m.Text == a.secret:
			a.allowedChats[m.Chat.Recipient()] = true
			log.WithField("component", "auth middleware").
				Infof("%v chat added to white list", m.Chat.Recipient())
			return
		case m.Text == "save all data":
			err := a.saveAuthToFile()
			if err != nil {
				log.WithField("component", "auth middleware").
					Error(err)
			}
			return
		case a.allowedChats[m.Chat.Recipient()]:
			log.WithField("component", "auth middleware").
				Infof("Message from allowed chat list %v", m.Chat.Recipient())
		default:
			log.WithField("component", "auth middleware").
				Infof("Chat %v is not allowed", m.Chat.Recipient())
			return
		}

		next(m)
	}
}

func logTextMessageMiddleware(next func(m *tb.Message)) func(m *tb.Message) {
	return func(m *tb.Message) {
		log.WithField("component", "log middleware").
			Infof("Message from %s. Text: %s\n", m.Sender.Username, m.Text)
		next(m)
	}
}
