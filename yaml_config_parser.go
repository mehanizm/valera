package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// config is struct to parse yaml configuration
type config struct {
	proxyURL   string `yaml:"PROXY_URL"`
	proxyUser  string `yaml:"PROXY_USER"`
	proxyPass  string `yaml:"PROXY_PASS"`
	tgBotToken string `yaml:"TELEGRAM_BOT_TOKEN"`
	jiraURL    string `yaml:"JIRA_URL"`
	jiraUser   string `yaml:"JIRA_USER"`
	jiraPass   string `yaml:"JIRA_PASS"`
}

// ParseFromFile parse YAML congig from file
// input – path to congif file
// structure of the file should corresponds to YamlConfig struct
func parseConfigFromFile(filePath string) (*config, error) {

	filename, _ := filepath.Abs(filePath)

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Cannot open configuration file", err)
		return nil, err
	}

	var c config
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println("Cannot parse yaml configuration", err)
		return nil, err
	}

	log.Println("Read YAML configuration file")
	return &c, nil

}
