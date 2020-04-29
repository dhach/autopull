package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"os"
)

var dockerCLI *client.Client
var localImages []types.ImageSummary

var log = logrus.New()
var ctx context.Context = context.Background()

func initDockerCLI() (cli *client.Client) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	cli.NegotiateAPIVersion(ctx)
	if checkError(err) {
		log.Fatal("Cannot create instance of Docker client")
	}
	return
}

func main() {
	configureLogger()
	log.Debug("Starting application")

	arg := os.Args[1:]
	if len(arg) != 1 {
		log.Fatal("Need exactly 1 argument, got: ", len(arg))
		os.Exit(99)
	}

	dockerCLI = initDockerCLI()
	localImages = getLocalImages()

	// configFilePath := "./local_test_config.json"
	configs := loadConfigs(arg[0])

	for _, config := range configs {
		runner := NewRunner(config)
		err := runner.Run()
		if checkError(err) {
			log.Fatal("Error executing command: ", err)
			os.Exit(1)
		}
	}
	log.Info("All done")
	os.Exit(0)
}
