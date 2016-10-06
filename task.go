package podio

import "encoding/json"
import "fmt"

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) GetTasksJson(params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
  fmt.Println("original params", params)
  params["space"] = 4412774
  fmt.Println("Sending params", params)
	err = client.RequestWithParams("GET", "/task/", nil, params, &rawResponse)
	return
}
