package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// config is struct to parse yaml configuration
type config struct {
	PROXY_URL          string `yaml:"PROXY_URL"`
	PROXY_USER         string `yaml:"PROXY_USER"`
	PROXY_PASS         string `yaml:"PROXY_PASS"`
	TELEGRAM_BOT_TOKEN string `yaml:"TELEGRAM_BOT_TOKEN"`
	JIRA_USER          string `yaml:"JIRA_USER"`
	JIRA_PASS          string `yaml:"JIRA_PASS"`
	JIRA_URL           string `yaml:"JIRA_URL"`
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
