package podio

import "fmt"

// Status are messages posted to the space stream
type Status struct {
	Id int64 `json:"status_id"`
}

// https://developers.podio.com/doc/status/add-new-status-message-22336
func (client *Client) StatusCreate(spaceId int64, params map[string]interface{}) (s Status, err error) {
	path := fmt.Sprintf("/status/space/%d/", spaceId)
	err = client.RequestWithParams("POST", path, nil, params, &s)
	return
}
