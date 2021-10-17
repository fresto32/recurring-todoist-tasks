package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type TaskJson struct {
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

type AddTaskInput struct {
	Project   string
	Content   string
	Labels    []string
	Priority  int
	DueString string
}

type addTaskRequest struct {
	ProjectId int    `json:"project_id"`
	Content   string `json:"content"`
	LabelIds  []int  `json:"labels_id"`
	Priority  int    `json:"priority"`
	DueString string `json:"due_string"`
}

type AddTaskResponse struct {
	Id           int    `json:"id"`
	Order        int    `json:"order"`
	Priority     int    `json:"priority"`
	ProjectId    int    `json:"project_id"`
	SectionId    int    `json:"section_id"`
	ParentId     int    `json:"parent_id"`
	Url          string `json:"url"`
	CommentCount int    `json:"comment_count"`
	Completed    bool   `json:"completed"`
	Content      string `json:"content"`
	Description  string `json:"description"`
	Due          struct {
		Date       string `json:"date"`
		DateTime   string `json:"datetime"`
		Recurring  bool   `json:"recurring"`
		DateString string `json:"string"`
		Timezome   string `json:"timezone"`
	} `json:"due"`
}

func AllTasks(apiToken string) chan []TaskJson {
	tasks := make(chan []TaskJson)

	go func() {
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

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var receivedTasks []TaskJson

		err = json.Unmarshal([]byte(body), &receivedTasks)
		if err != nil {
			panic(err)
		}

		tasks <- receivedTasks

		close(tasks)
	}()

	return tasks
}

func AddTask(apiToken string, addTaskInput AddTaskInput) chan AddTaskResponse {
	response := make(chan AddTaskResponse)

	go func() {
		projectChan := FindProjectByName(apiToken, addTaskInput.Project)
		labelChans := make([]chan LabelJson, len(addTaskInput.Labels))

		for i, v := range addTaskInput.Labels {
			labelChans[i] = FindLabelByName(apiToken, v)
		}

		labelIds := make([]int, len(addTaskInput.Labels))

		for i, v := range labelChans {
			labelIds[i] = (<-v).Id
		}

		addTaskReq := &addTaskRequest{
			ProjectId: (<-projectChan).Id,
			Content:   addTaskInput.Content,
			LabelIds:  labelIds,
			Priority:  addTaskInput.Priority,
			DueString: addTaskInput.DueString,
		}

		addTaskReqJson, err := json.Marshal(addTaskReq)
		if err != nil {
			panic(err)
		}

		req, err := http.NewRequest("POST", "https://api.todoist.com/rest/v1/tasks", bytes.NewBuffer((addTaskReqJson)))
		if err != nil {
			panic(err)
		}

		req.Header = http.Header{
			"Content-Type":  []string{"application/json"},
			"Authorization": []string{"Bearer " + apiToken},
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		if resp.StatusCode != 200 {
			panic("Status code is not 200")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var addTaskResponse AddTaskResponse
		err = json.Unmarshal([]byte(body), &addTaskResponse)
		if err != nil {
			panic(err)
		}

		response <- addTaskResponse

		close(response)
	}()

	return response
}
