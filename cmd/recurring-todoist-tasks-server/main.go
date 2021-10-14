package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	apiToken := "some token"

	projects := make(chan string)
	go printProjects(apiToken, projects)

	for p := range projects {
		log.Printf(p)
	}
}

func printProjects(apiToken string, projects chan string) {
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
		//Handle Error
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//Convert the body to type string
	sb := string(body)

	projects <- sb
}
