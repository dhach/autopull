package main

import (
	"bytes"
	"github.com/docker/docker/api/types"
	"regexp"
)

func pullImage(name string) (changed bool, pullError error) {
	changed = false
	re := regexp.MustCompile(`.*Status: Image is up to date.*`)

	reader, pullError := dockerCLI.ImagePull(ctx, name, types.ImagePullOptions{})
	if checkError(pullError) {
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	isUpToDate := re.Match(buf.Bytes())
	if isUpToDate == false {
		changed = true
	}
	defer reader.Close()

	return
}

func getLocalImages() (localImages []types.ImageSummary) {
	localImages, err := dockerCLI.ImageList(ctx, types.ImageListOptions{})
	if checkError(err) {
		log.Fatal("Cannot list local images: ", err)
	}
	return
}
