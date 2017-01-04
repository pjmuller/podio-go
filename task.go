package podio

import "encoding/json"

type Task struct {
	Id   		int64  	`json:"task_id"`
	Status 	string 	`json:"status"`
	Text 		string 	`json:"text"`
	Description string `json:"description"`
	DueOn 	string  `json:"due_on"` // we pick string as sometimes blank
	Ref 		TaskRef `json:"ref"`
	SpaceId int     `json:"space_id"`
}

type TaskRef struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasksJson(params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	err = client.RequestWithParams("GET", "/task/", nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasks(params map[string]interface{}) (tasks []Task, err error) {
	err = client.RequestWithParams("GET", "/task/", nil, params, &tasks)
	return
}
