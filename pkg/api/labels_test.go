package api_test

import (
	"os"
	"testing"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func TestAllLabels(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	if <-api.AllLabels(apiToken) == nil {
		t.Fatalf("No labels found")
	}
}

func TestFindLabelByName(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	label := <-api.FindLabelByName(apiToken, "home")

	if label.Id != 2154881078 {
		t.Fatalf(`GetLabelOfName(...) = %v, expected ID of %q`, label, 1)
	}
}
