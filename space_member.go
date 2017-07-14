package podio

import "fmt"

type SpaceMember struct {
	Profile Contact `json:"profile"`
	Role    string  `json:"role"`
}

// https://developers.podio.com/doc/space-members/get-space-members-v2-19350328
func (client *Client) FindAllForSpace(id int64, options map[string]interface{}) (spaceMembers []SpaceMember, err error) {
	path := fmt.Sprintf("/space/%d/member/v2/", id)
	err = client.RequestWithParams("GET", path, nil, options, &spaceMembers)
	return
}
