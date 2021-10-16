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

func GetAllLabels(apiToken string, labels chan []LabelJson) {
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

	var receivedLabels []LabelJson

	err = json.Unmarshal([]byte(body), &receivedLabels)
	if err != nil {
		panic(err)
	}

	labels <- receivedLabels

	close(labels)
}

func GetLabelOfName(apiToken string, name string) chan LabelJson {
	label := make(chan LabelJson)

	go func() {
		allLabelsChannel := make(chan []LabelJson)
		go GetAllLabels(apiToken, allLabelsChannel)

		allLabels := <-allLabelsChannel

		for _, v := range allLabels {
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
