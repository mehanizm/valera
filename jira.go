package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func (b *bot) getIssueDescFromJira(issueKey string) string {

	issue, _, err := b.Issue.Get(issueKey, nil)
	if err != nil {
		log.WithField("component", "jira issue parser").Error(err)
		return ""
	}

	return fmt.Sprintf("%s: %+v\n", issue.Key, issue.Fields.Summary)

}
