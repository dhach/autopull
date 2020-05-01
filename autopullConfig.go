package main

import (
	"encoding/json"
	"io/ioutil"
)

type jsonConfig struct {
	// Image is a simple string which points to the image that should be checked
	Image string `json:"image"`
	// Tag is a simple string holding the image tag to be checked. Defaults to "latest"
	Tag string `json:"tag"`
	// Action specifies the commands which should be executed upon a successfull pull
	Actions []string `json:"actions"`
}

func loadConfigs(configFilePath string) (configList []jsonConfig) {
	configFileContent, readErr := ioutil.ReadFile(configFilePath)
	if checkError(readErr) {
		log.Fatal("Cannot read configuration file: ", readErr)
	}

	jsonErr := json.Unmarshal(configFileContent, &configList)
	if checkError(jsonErr) {
		log.Fatal("Cannot unmarshal JSON data: ", jsonErr)
	}

	return
}
