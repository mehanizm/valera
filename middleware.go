package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

type authData struct {
	secret       string
	allowedChats map[string]bool
	filePath     string
}

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

	log.Println("Data saved to file", a.filePath)

	return nil
}

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

	log.Printf("Successfully read %v lines from auth file\n", len(a.allowedChats))

	return nil
}

func (a *authData) checkAuthMiddleware(next func(m *tb.Message)) func(m *tb.Message) {
	return func(m *tb.Message) {

		switch {
		case m.Text == a.secret:
			a.allowedChats[m.Chat.Recipient()] = true
			log.Printf("%v chat added to white list\n", m.Chat.Recipient())
			return
		case m.Text == "save all data":
			err := a.saveAuthToFile()
			if err != nil {
				log.Println("ERROR:", err)
			}
			return
		case a.allowedChats[m.Chat.Recipient()]:
			log.Printf("Message from allowed chat list %v\n", m.Chat.Recipient())
		default:
			log.Printf("Chat %v is not allowed\n", m.Chat.Recipient())
			return
		}

		next(m)
	}
}

func logTextMessageMiddleware(next func(m *tb.Message)) func(m *tb.Message) {
	return func(m *tb.Message) {
		log.Printf("Message from %s. Text: %s\n", m.Sender.Username, m.Text)
		next(m)
	}
}
