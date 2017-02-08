package podio

import "fmt"

type SpaceMember struct {
	Profile  Contact  `json:"profile"`
	Role     string `json:"role"`
}

func (client *Client) FindAllForSpace(id int64) (spaceMembers []SpaceMember, err error) {
	path := fmt.Sprintf("/space/%d/member/v2/", id)
	err = client.Request("GET", path, nil, nil, &spaceMembers)
	return
}
