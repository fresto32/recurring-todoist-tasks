package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Task struct {
	Id            int    `json:"id"`
	Assigner      int    `json:"assigner"`
	Project_id    int    `json:"project_id"`
	Section_id    int    `json:"section_id"`
	Order         int    `json:"order"`
	Content       string `json:"content"`
	Description   string `json:"description"`
	Completed     bool   `json:"completed"`
	Label_ids     []int  `json:"label_ids"`
	Priority      int    `json:"priority"`
	Comment_count int    `json:"comment_count"`
	Creator       int    `json:"creator"`
	Created       string `json:"created"`
	Url           string `json:"url"`
	Due           struct {
		Recurring bool   `json:"recurring"`
		String    string `json:"string"`
		Date      string `json:"date"`
	} `json:"due"`
}

func GetTasks(apiToken string, tasks chan []Task) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v1/tasks", nil)
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

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var receivedTasks []Task

	err = json.Unmarshal([]byte(body), &receivedTasks)
	if err != nil {
		panic(err)
	}

	tasks <- receivedTasks

	close(tasks)
}
