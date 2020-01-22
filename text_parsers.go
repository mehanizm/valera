package main

import (
	"regexp"
)

var regTemplate, _ = regexp.Compile("[A-Za-z]{1,9}-[0-9]{1,5}")

func parseIssueKeysFromMsg(msg string) []string {
	return getUniqueSlice(regTemplate.FindAllString(msg, -1))
}
