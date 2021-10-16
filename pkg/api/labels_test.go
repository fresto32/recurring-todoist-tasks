package api_test

import (
	"os"
	"testing"

	"github.com/fresto32/recurring-todoist-tasks/pkg/api"
)

func TestGetAllLabels(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	labelsChannel := make(chan []api.LabelJson)
	go api.GetAllLabels(apiToken, labelsChannel)

	labels := <-labelsChannel

	if labels == nil {
		t.Fatalf("No labels found")
	}
}

func TestGetLabelOfName(t *testing.T) {
	apiToken := os.Getenv("API_TOKEN")

	labelCh := api.GetLabelOfName(apiToken, "home")

	label := <-labelCh

	if label.Id != 2154881078 {
		t.Fatalf(`GetLabelOfName(...) = %v, expected ID of %q`, label, 1)
	}
}
