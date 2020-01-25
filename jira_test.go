package main

import (
	"testing"

	jira "github.com/andygrunwald/go-jira"
)

func Test_printToMessage(t *testing.T) {
	tests := []struct {
		name string
		issue *jira.Issue
		want string
	}{
		{
			name: "Case 1. Positive",
issue: &jira.Issue{
	Key: "CCC-1",
	Fields: &jira.IssueFields {
		Summary: "Test issue",
	},
},
want: "CCC-1: Test issue\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printToMessage(tt.issue); got != tt.want {
				t.Errorf("printToMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
