package podio

import "fmt"

type SpaceMember struct {
	Profile Contact `json:"profile"`
	Role    string  `json:"role"`
}

type SpaceMemberV1 struct {
	Profile ContactSimple `json:"profile"`
	User    UserSimple    `json:"user"`
	Role    string        `json:"role"`
}

// Contact describes a Podio contact object
type ContactSimple struct {
	ProfileId int    `json:"profile_id"`
	Name      string `json:"name"`
}

// https://developers.podio.com/doc/space-members/get-space-members-v2-19350328
func (client *Client) FindAllForSpace(id int64, options map[string]interface{}) (spaceMembers []SpaceMember, err error) {
	path := fmt.Sprintf("/space/%d/member/v2/", id)
	err = client.RequestWithParams("GET", path, nil, options, &spaceMembers)
	return
}

// https://developers.podio.com/doc/space-members/get-members-of-space-22395
func (client *Client) FindAllForSpaceV1(id int64, options map[string]interface{}) (spaceMembers []SpaceMemberV1, err error) {
	path := fmt.Sprintf("/space/%d/member/", id)
	err = client.RequestWithParams("GET", path, nil, options, &spaceMembers)
	return
}
