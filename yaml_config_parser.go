package main

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// config is struct to parse and store yaml configuration
type config struct {
	ProxyURL   string `yaml:"PROXY_URL"`
	ProxyUser  string `yaml:"PROXY_USER"`
	ProxyPass  string `yaml:"PROXY_PASS"`
	TgBotToken string `yaml:"TELEGRAM_BOT_TOKEN"`
	JiraURL    string `yaml:"JIRA_URL"`
	JiraUser   string `yaml:"JIRA_USER"`
	JiraPass   string `yaml:"JIRA_PASS"`
}

// parseConfigFromFile parse YAML congig from file
// filePath – path to congif file
// structure of the file should corresponds to YamlConfig struct
func parseConfigFromFile(filePath string) (*config, error) {

	filename, err := filepath.Abs(filePath)
	if err != nil {
		log.WithField("component", "yaml file open").Error(err)
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.WithField("component", "yaml file read").Error(err)
		return nil, err
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.WithField("component", "yaml file parse").Error(err)
		return nil, err
	}

	log.WithField("component", "yaml parser").Info("configuration parsed successfully")
	return &c, nil

}
