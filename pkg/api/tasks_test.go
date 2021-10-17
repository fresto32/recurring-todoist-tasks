package api_test

import (
	"os"
	"testing"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func TestAllTasks(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	if <-api.AllTasks(apiToken) == nil {
		t.Fatalf("No labels found")
	}
}

func TestAddTask(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	testTask := api.AddTaskInput{
		Project:   "One Off Tasks",
		Content:   "Test task",
		Labels:    []string{"home", "next"},
		Priority:  4,
		DueString: "today",
	}

	response := <-api.AddTask(apiToken, testTask)

	if response.Id == 0 {
		t.Fatalf("Could not add task")
	}
}
