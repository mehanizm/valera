package main

import (
	"fmt"
	"log"
)

func (b *bot) getIssueDescFromJira(issueKey string) string {

	issue, _, err := b.Issue.Get(issueKey, nil)
	if err != nil {
		log.Println("ERROR in jira: ", err)
		return ""
	}

	return fmt.Sprintf("%s: %+v\n", issue.Key, issue.Fields.Summary)

}
