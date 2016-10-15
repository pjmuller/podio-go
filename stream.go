package podio

import (
  "encoding/json"
  "fmt"
)

// https://developers.podio.com/doc/tasks/get-tasks-77949
func (client *Client) StreamForSpaceV3(spaceId int64, params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
  path := fmt.Sprintf("/stream/space/%d/v3/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &rawResponse)
	return
}
