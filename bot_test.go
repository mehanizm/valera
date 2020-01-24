package main

import (
	"reflect"
	"testing"
)

func Test_initializeBot(t *testing.T) {
	type args struct {
		config *config
	}
	tests := []struct {
		name    string
		config  *config
		want    *bot
		wantErr bool
	}{
		{
			name:    "CASE 1. Empty config",
			config:  nil,
			want:    nil,
			wantErr: true,
		},
		{
			name: "CASE 2. Not existing proxy URL",
			config: &config{
				ProxyURL:  "test-example.ru",
				ProxyUser: "test",
				ProxyPass: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CASE 3. Not existing jira URL",
			config: &config{
				JiraURL:    "test-example.ru",
				JiraUser:   "test",
				JiraPass:   "test",
				TgBotToken: "709294056:AAFIfAgcNTzaIjPdR5bColgOj2vGOKSsgTg",
				ProxyURL:   "grsst.s5.opennetwork.cc:999",
				ProxyUser:  "41591017",
				ProxyPass:  "5NEISabl",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CASE 4. Not valid bot token",
			config: &config{
				TgBotToken: "123",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "CASE 5. Not reachable telegram api",
			config: &config{
				TgBotToken: "709294056:AAFIfAgcNTzaIjPdR5bColgOj2vGOKSsgTg",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initializeBot(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("initializeBot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeBot() = %v, want %v", got, tt.want)
			}
		})
	}
}
