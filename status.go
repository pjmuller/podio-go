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

// https://developers.podio.com/doc/status/update-a-status-message-22338
func (client *Client) StatusUpdate(statusID int64, params map[string]interface{}) error {
	path := fmt.Sprintf("/status/%d/", statusID)
	return client.RequestWithParams("PUT", path, nil, params, nil)
}

// https://developers.podio.com/doc/status/delete-a-status-message-22339
func (client *Client) StatusDelete(statusID int64) error {
	path := fmt.Sprintf("/status/%d/", statusID)
	return client.Request("DELETE", path, nil, nil, nil)
}
