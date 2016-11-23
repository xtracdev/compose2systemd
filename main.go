package main

import (
	"flag"
	"github.com/docker/libcompose/project"
	"log"
	"github.com/docker/libcompose/config"
	"fmt"
	"text/template"
	"os"
)

var (
	composeFile string
	appName string
	unitFileTemplate *template.Template
)

func init() {
	flag.StringVar(&composeFile, "compose-file-path", "./", "Specify an alternate path for compose files")
	flag.StringVar(&appName, "app-name", "", "Application name described by docker compose file")

	var err error
	unitFileTemplate = template.New("unit")
	unitFileTemplate = unitFileTemplate.Delims("[[","]]")
	unitFileTemplate, err = unitFileTemplate.Parse(UnitTemplate)
	if err != nil {
		log.Fatalf("Unabled to parse unit template: %s", err.Error())
	}
}

func main() {

	flag.Parse()
	if appName == "" {
		log.Fatal("Specify app-name on the command line")
	}

	dockerCompose := project.NewProject(&project.Context{
		ProjectName:  "kube",
		ComposeFiles: []string{composeFile},
	}, nil, &config.ParseOptions{})


	if err := dockerCompose.Parse(); err != nil {
		log.Fatalf("Failed to parse the compose project from %s: %v", composeFile, err)
	}

	for _, name := range dockerCompose.ServiceConfigs.Keys() {
		log.Printf("Generating unit for service %s",name)

		service, ok := dockerCompose.GetServiceConfig(name)
		if !ok {
			log.Println("WARNING: Unable to get service config for %s... skipping.")
			continue
		}

		dependency := "docker"
		if len(service.DependsOn) > 0 {
			if len(service.DependsOn) > 1 {
				log.Println("WARNING - multiple dependencies exist - using the first dependency only")
			}
			dependency = service.DependsOn[0]
		}

		unit := Unit {
			Service:name,
			Dependency:dependency,
			AppName:appName,
		}



		fmt.Println("")
		err := unitFileTemplate.Execute(os.Stdout, unit)
		if err != nil {
			log.Fatalf("Error executing template: %s", err.Error())
		}
		fmt.Println("")
	}
}
