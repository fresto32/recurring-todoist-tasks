package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Label struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Order    int    `json:"order"`
	Color    int    `json:"color"`
	Favorite bool   `json:"favorite"`
}

func GetLabels(apiToken string, labels chan []Label) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v1/labels", nil)
	if err != nil {
		panic(err)
	}

	req.Header = http.Header{
		"Authorization": []string{"Bearer " + apiToken},
	}

	resp, err := client.Do(req)
	if err != nil {
		//Handle Error
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var receivedLabels []Label

	err = json.Unmarshal([]byte(body), &receivedLabels)
	if err != nil {
		panic(err)
	}

	labels <- receivedLabels

	close(labels)
}
