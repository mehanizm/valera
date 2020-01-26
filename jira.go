package main

import (
	"fmt"

	jira "github.com/andygrunwald/go-jira"
	log "github.com/sirupsen/logrus"
)

// getIssueDescFromJira
// get information from jira about
// issue by key
func (b *bot) getIssueDescFromJira(issueKey string) string {

	issue, _, err := b.Issue.Get(issueKey, nil)
	if err != nil {
		log.WithField("component", "jira issue parser").Error(err)
		return ""
	}

	log.WithField("component", "jira issue parser").
		Infof("get result from jira about task %v", issueKey)

	return printToMessage(issue)

}

func printToMessage(issue *jira.Issue) string {
	return fmt.Sprintf("%s: %s\n", issue.Key, issue.Fields.Summary)
}
