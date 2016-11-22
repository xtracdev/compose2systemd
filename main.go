package main

import (
	"flag"
	"github.com/docker/libcompose/project"
	"log"
	"github.com/docker/libcompose/config"
)

var (
	composeFile string
)

func init() {
	flag.StringVar(&composeFile, "compose-file-path", "./", "Specify an alternate path for compose files")
}

func main() {

	flag.Parse()

	dockerCompose := project.NewProject(&project.Context{
		ProjectName:  "kube",
		ComposeFiles: []string{composeFile},
	}, nil, &config.ParseOptions{})


	if err := dockerCompose.Parse(); err != nil {
		log.Fatalf("Failed to parse the compose project from %s: %v", composeFile, err)
	}

	for _, name := range dockerCompose.ServiceConfigs.Keys() {
		log.Printf("%s",name)
	}
}
