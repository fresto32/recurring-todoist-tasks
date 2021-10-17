package main

import (
	"log"
	"os"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")

	projects := <-api.AllProjects(apiToken)
	for p := range projects {
		log.Printf("%v\n", p)
	}

	tasks := <-api.AllTasks(apiToken)
	for t := range tasks {
		log.Printf("%v\n", t)
	}

	labels := <-api.AllLabels(apiToken)
	for l := range labels {
		log.Printf("%v\n", l)
	}
}
