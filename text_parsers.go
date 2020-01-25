package main

import (
	"regexp"
)

var regTemplate, _ = regexp.Compile("[A-Za-z]{1,9}-[0-9]{1,5}")

// parseIssueKeysFromMsg
// parse jira issue keys from plain text
// telegram chat message
func parseIssueKeysFromMsg(msg string) []string {
	if msg == "" {
		return make([]string, 0)
	}
	return getUniqueSlice(regTemplate.FindAllString(msg, -1))
}
