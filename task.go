package podio

import (
	"encoding/json"
	"fmt"
)

type Task struct {
	Id          int64   `json:"task_id"`
	Status      string  `json:"status"`
	Text        string  `json:"text"`
	Description string  `json:"description"`
	DueOn       *string `json:"due_on"` // we pick string as sometimes blank

	CompletedOn *string `json:"completed_on"`
	CompletedBy TaskRef `json:"completed_by"`
	DeletedOn   *string `json:"deleted_on"`
	DeletedBy   TaskRef `json:"deleted_by"`
	CreatedOn   *string `json:"created_on"`
	CreatedBy   TaskRef `json:"created_by"`
	Responsible Contact `json:"responsible"`

	Ref     TaskRef `json:"ref"`
	Private bool    `json:"private"`

	SpaceId    int             `json:"space_id"`
	ExternalId int64           `json:"external_id"`
	Labels     []*TaskLabel    `json:"labels"`
	Recurrence json.RawMessage `json:"recurrence"`
	Reminder   TaskReminder    `json:"reminder"`
}

type TaskRef struct {
	Id   int64    `json:"id"`
	Type string   `json:"type"`
	Data TaskData `json:"data"`
}

type TaskData struct {
	App TaskApp `json:"app"`
}

type TaskApp struct {
	Id int64 `json:"app_id"`
}

type TaskLabel struct {
	Id    int    `json:"label_id"`
	Text  string `json:"text"`
	Color string `json:"color"`
}

type TaskLabelSimple struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}

type TaskReminder struct {
	Delta *int `json:"remind_delta"`
}

type TaskCount struct {
	Count int `json:"count"`
}

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasksJson(params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	err = client.RequestWithParams("GET", "/task/", nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/tasks/get-task-22413
func (client *Client) GetTask(taskID int64) (task Task, err error) {
	path := fmt.Sprintf("/task/%d", taskID)
	err = client.Request("GET", path, nil, nil, &task)
	return
}

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasks(params map[string]interface{}) (tasks []Task, err error) {
	err = client.RequestWithParams("GET", "/task/", nil, params, &tasks)
	return
}

// https://developers.podio.com/doc/tasks/get-task-count-38316458
func (client *Client) GetTaskCount(refType string, refId int64) (count TaskCount, err error) {
	path := fmt.Sprintf("/task/%s/%d/count", refType, refId)
	err = client.Request("GET", path, nil, nil, &count)
	return
}
