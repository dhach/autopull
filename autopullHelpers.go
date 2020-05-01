package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

// checkError checks if the passed error is nil. It returns true if there is an error and false if not
func checkError(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func concatImageName(name string, tag string) (result string) {
	var imageNameStringBuilder strings.Builder

	imageNameStringBuilder.WriteString(name)
	imageNameStringBuilder.WriteString(":")
	imageNameStringBuilder.WriteString(tag)

	result = imageNameStringBuilder.String()

	return
}

func configureLogger() {
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetReportCaller(true)
	log.Out = os.Stdout

	level := os.Getenv("LOGLEVEL")
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}
