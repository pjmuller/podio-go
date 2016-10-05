package podio

import "fmt"

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasksJson(params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	err = client.RequestWithParams("GET", "/task/", nil, params, &rawResponse)
	return
}
