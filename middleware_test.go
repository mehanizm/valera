package main

import (
	"os"
	"reflect"
	"testing"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Test_authData_saveAuthToFile(t *testing.T) {
	type fields struct {
		secret       string
		allowedChats map[string]bool
		filePath     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Case 1. Print file",
			fields: fields{
				filePath: "test/white_list_1.txt",
				allowedChats: map[string]bool{
					"test_1": true,
					"test_2": true,
					"test_3": true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authData{
				secret:       tt.fields.secret,
				allowedChats: tt.fields.allowedChats,
				filePath:     tt.fields.filePath,
			}
			if err := a.saveAuthToFile(); (err != nil) != tt.wantErr {
				t.Errorf("authData.saveAuthToFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			a.allowedChats = make(map[string]bool, 0)
			a.readAuthFromFile()
			eq := reflect.DeepEqual(tt.fields.allowedChats, a.allowedChats)
			if !eq {
				t.Errorf("Mistake in read and write from auth file")
			}
			os.Remove(tt.fields.filePath)
		})
	}
}

func Test_logTextMessageMiddleware(t *testing.T) {
	tests := []struct {
		name string
		next func(m *tb.Message)
		want func(m *tb.Message)
	}{
		{
			name: "Case 1. Check positive",
			next: func(m *tb.Message) {},
			want: func(m *tb.Message) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if reflect.TypeOf(logTextMessageMiddleware(tt.next)) != reflect.TypeOf(tt.want) {
				t.Error("Error in logTextMessageMiddleware")
			}
		})
	}
}

func Test_authData_checkAuthMiddleware(t *testing.T) {
	type fields struct {
		secret       string
		allowedChats map[string]bool
		filePath     string
	}
	tests := []struct {
		name   string
		fields fields
		next   func(m *tb.Message)
		want   func(m *tb.Message)
	}{
		{
			name: "Case 1. Positive",
			fields: fields{
				allowedChats: map[string]bool{
					"123": true,
				},
			},
			next: func(m *tb.Message) {},
			want: func(m *tb.Message) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authData{
				allowedChats: tt.fields.allowedChats,
			}
			if reflect.TypeOf(a.checkAuthMiddleware(tt.next)) != reflect.TypeOf(tt.want) {
				t.Error("Error in logTextMessageMiddleware")
			}
		})
	}
}
