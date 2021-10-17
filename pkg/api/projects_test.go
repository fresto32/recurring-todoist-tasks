package api_test

import (
	"os"
	"testing"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func TestAllProjects(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	if <-api.AllProjects(apiToken) == nil {
		t.Fatalf("No projects found")
	}
}

func TestFindProjectByName(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	project := <-api.FindProjectByName(apiToken, "One Off Tasks")

	if project.Id != 2240138508 {
		t.Fatalf(`GetProjectOfName(...) = %v, expected ID of %q`, project, 1)
	}
}
