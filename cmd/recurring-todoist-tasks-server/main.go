package main

import (
	"log"
	"os"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func main() {
	apiToken := os.Getenv("API_TOKEN")

	projects := make(chan []api.ProjectJson)
	go api.GetAllProjects(apiToken, projects)

	for p := range projects {
		log.Printf("%v\n", p)
	}

	tasks := make(chan []api.TaskJson)
	go api.GetAllTasks(apiToken, tasks)

	for t := range tasks {
		log.Printf("%v\n", t)
	}

	labels := make(chan []api.LabelJson)
	go api.GetAllLabels(apiToken, labels)

	for l := range labels {
		log.Printf("%v\n", l)
	}
}
