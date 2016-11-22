package main

import (
	"flag"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
	"github.com/docker/libcompose/docker"
	"log"
)

var (
	composeFile string
)

func init() {
	flag.StringVar(&composeFile, "compose-file-path", "./", "Specify an alternate path for compose files")
}

func main() {
	flag.Parse()

	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{composeFile},
			ProjectName:  "my-compose",
		},
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	if err := project.Parse(); err != nil {
		log.Fatalf("Failed to parse the compose project from %s: %v", composeFile, err)
	}
}
