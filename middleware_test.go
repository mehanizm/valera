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
	testCase := struct {
		name string
		next func(m *tb.Message)
	}{
		name: "Case 1. Check positive",
		next: func(m *tb.Message) {},
	}

	t.Run(testCase.name, func(t *testing.T) {
		if reflect.TypeOf(logTextMessageMiddleware(testCase.next)) != reflect.TypeOf(testCase.next) {
			t.Error("Error in logTextMessageMiddleware")
		}
	})
}

func Test_authData_checkAuthMiddleware_positive(t *testing.T) {
	testCase := struct {
		name   string
		fields *authData
		next   func(m *tb.Message)
		want   func(m *tb.Message)
	}{
		name: "Case 1. Check positive",
		fields: &authData{
			allowedChats: map[string]bool{
				"123": true,
			},
		},
		next: func(m *tb.Message) {},
		want: func(m *tb.Message) {},
	}

	t.Run(testCase.name, func(t *testing.T) {
		if reflect.TypeOf(testCase.fields.checkAuthMiddleware(testCase.next)) != reflect.TypeOf(testCase.want) {
			t.Error("Error in logTextMessageMiddleware")
		}
	})

}

func assertPanic(t *testing.T, f func(m *tb.Message), m *tb.Message) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code not panic")
		}
	}()
	f(m)
}

func Test_authData_checkAuthMiddleware(t *testing.T) {
	tests := []struct {
		name      string
		auth      *authData
		m         *tb.Message
		wantPanic bool
	}{
		{
			name: "Case 1. Check not allowed",
			auth: &authData{
				allowedChats: map[string]bool{
					"123": true,
				},
				// secret: "secret",
			},
			m: &tb.Message{
				Text: "some text",
				Chat: &tb.Chat{
					ID: 122,
				},
			},
			wantPanic: false,
		},
		{
			name: "Case 2. Check allowed",
			auth: &authData{
				allowedChats: map[string]bool{
					"123": true,
				},
				// secret: "secret",
			},
			m: &tb.Message{
				Text: "some text",
				Chat: &tb.Chat{
					ID: 123,
				},
			},
			wantPanic: true,
		},
		{
			name: "Case 3. Check secret",
			auth: &authData{
				allowedChats: map[string]bool{
					"123": true,
				},
				secret: "secret",
			},
			m: &tb.Message{
				Text: "secret",
				Chat: &tb.Chat{
					ID: 122,
				},
			},
			wantPanic: false,
		},
		{
			name: "Case 4. Check save",
			auth: &authData{
				allowedChats: map[string]bool{
					"123": true,
				},
				secret: "secret",
			},
			m: &tb.Message{
				Text: "save all data",
				Chat: &tb.Chat{
					ID: 123,
				},
			},
			wantPanic: false,
		},
		{
			name: "Case 5. Check save not allowed",
			auth: &authData{
				allowedChats: map[string]bool{
					"123": true,
				},
				secret: "secret",
			},
			m: &tb.Message{
				Text: "save all data",
				Chat: &tb.Chat{
					ID: 122,
				},
			},
			wantPanic: false,
		},
	}
	next := func(m *tb.Message) {
		panic("panic")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			afterMiddleware := tt.auth.checkAuthMiddleware(next)
			defer func() {
				if r := recover(); r == nil && tt.wantPanic {
					t.Errorf("Auth middleware error")
				}
			}()
			afterMiddleware(tt.m)
		})
	}
}
