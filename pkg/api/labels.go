package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LabelJson struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Order    int    `json:"order"`
	Color    int    `json:"color"`
	Favorite bool   `json:"favorite"`
}

var allLabels []LabelJson

func AllLabels(apiToken string) chan []LabelJson {
	labels := make(chan []LabelJson)

	go func() {
		if allLabels != nil {
			labels <- allLabels
			return
		}

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
			panic(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var receivedLabels []LabelJson

		err = json.Unmarshal([]byte(body), &receivedLabels)
		if err != nil {
			panic(err)
		}

		labels <- receivedLabels

		allLabels = receivedLabels

		close(labels)
	}()

	return labels
}

func FindLabelByName(apiToken string, name string) chan LabelJson {
	label := make(chan LabelJson)

	go func() {
		for _, v := range <-AllLabels(apiToken) {
			if v.Name == name {
				label <- v
				close(label)
				return
			}
		}
		close(label)
	}()

	return label
}
