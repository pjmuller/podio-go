package podio

import "fmt"

type GrantResponse struct {
	Invitable []Contact `json:"invitable"`
}

// https://developers.podio.com/doc/grants/create-grant-16168841
func (client *Client) CreateGrant(refType string, refId int64, params map[string]interface{}) (g GrantResponse, err error) {
	path := fmt.Sprintf("/grant/%s/%d", refType, refId)
	err = client.RequestWithParams("POST", path, nil, params, &g)
	return
}
