package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ProjectJson struct {
	Id            int    `json:"id"`
	Parent        int    `json:"parent"`
	Parent_id     int    `json:"parent_id"`
	Order         int    `json:"order"`
	Color         int    `json:"color"`
	Name          string `json:"name"`
	Comment_count int    `json:"comment_count"`
	Shared        bool   `json:"shared"`
	Favorite      bool   `json:"favorite"`
	Inbox_project bool   `json:"inbox_project"`
	Sync_id       int    `json:"sync_id"`
	Url           string `json:"url"`
}

func GetAllProjects(apiToken string, projects chan []ProjectJson) {
	client := http.Client{}

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v1/projects", nil)
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

	var receivedProjects []ProjectJson

	err = json.Unmarshal([]byte(body), &receivedProjects)
	if err != nil {
		panic(err)
	}

	projects <- receivedProjects

	close(projects)
}

func GetProjectOfName(apiToken string, name string) chan ProjectJson {
	project := make(chan ProjectJson)

	go func() {
		allProjectsChannel := make(chan []ProjectJson)
		go GetAllProjects(apiToken, allProjectsChannel)

		allProjects := <-allProjectsChannel

		for _, v := range allProjects {
			if v.Name == name {
				project <- v
				close(project)
				return
			}
		}

		close(project)
	}()

	return project
}
