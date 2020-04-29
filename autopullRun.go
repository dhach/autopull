package main

import (
	"os/exec"
	"strings"
)

type autopullRunner struct {
	options jsonConfig
}

// NewRunner returns a new instance of an autopullRunner
func NewRunner(config jsonConfig) *autopullRunner {
	var runner autopullRunner
	runner.options = config
	return &runner
}

func (ar *autopullRunner) checkImageUpdate() (changed bool, imageDigestError error) {
	changed = false
	imageDigestError = nil
	imagePresent := false

	imageNameConstructed := concatImageName(ar.options.Image, ar.options.Tag)
	log.Debug("Starting check for: ", imageNameConstructed)

	for _, img := range localImages {

		imageInformation, _, inspectError := dockerCLI.ImageInspectWithRaw(ctx, img.ID)
		if checkError(inspectError) {
			imageDigestError = inspectError
			return
		}
		for _, tag := range imageInformation.RepoTags {
			if tag == imageNameConstructed {
				imagePresent = true
				hasChanged, pullError := pullImage(imageNameConstructed)
				if checkError(pullError) {
					imageDigestError = pullError
					return
				}
				if hasChanged {
					log.Info("Newer version of ", imageNameConstructed, " found")
					changed = true
				} else {
					log.Info("No newer version of ", imageNameConstructed, " available")
				}
			}
		}
	}
	if imagePresent == false {
		log.Warn("Image ", ar.options.Image, ":", ar.options.Tag, " not present on Docker host. Skipping...")
	}
	return
}

func (ar *autopullRunner) executeCommand() (execError error) {
	execError = nil

	commandList := strings.Split(ar.options.Action, " ")

	command := exec.Command(commandList[0])
	command.Args = commandList
	log.Info("Running command: ", commandList)
	out, execError := command.CombinedOutput()

	log.Debug("Stdout of command: ", string(out))
	log.Error("Stderr of command: ", execError)

	return
}

// Run instantiates clients for all defined Image/Tag/Action combinations
func (ar *autopullRunner) Run() (runError error) {
	runError = nil

	changed, err := ar.checkImageUpdate()
	if checkError(err) {
		runError = err
		return
	}
	if changed {
		err := ar.executeCommand()
		if checkError(err) {
			runError = err
			return
		}
	}

	return
}
