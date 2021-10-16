package api_test

import (
	"os"
	"testing"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func TestGetAllProjects(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	projectsChannel := make(chan []api.ProjectJson)
	go api.GetAllProjects(apiToken, projectsChannel)

	projects := <-projectsChannel

	if projects == nil {
		t.Fatalf("No projects found")
	}
}

func TestGetProjectOfName(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	project := <-api.GetProjectOfName(apiToken, "One Off Tasks")

	if project.Id != 2240138508 {
		t.Fatalf(`GetProjectOfName(...) = %v, expected ID of %q`, project, 1)
	}
}
