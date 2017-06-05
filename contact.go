package podio

import "fmt"

// Contact describes a Podio contact object
type Contact struct {
	UserId     int      `json:"user_id"`
	SpaceId    int      `json:"space_id"`
	Type       string   `json:"type"`
	Image      File     `json:"image"`
	ProfileId  int      `json:"profile_id"`
	OrgId      int      `json:"org_id"`
	Link       string   `json:"link"`
	Avatar     int      `json:"avatar"`
	LastSeenOn *Time    `json:"last_seen_on"`
	Name       string   `json:"name"`
	Emails     []string `json:"mail"`
}

const (
	maxPullContacts = 500
)

// https://developers.podio.com/doc/contacts/get-contacts-22400
func (client *Client) GetContacts(limit, offset int) (contacts []Contact, err error) {
	if limit == 0 {
		limit = maxPullContacts
	}
	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"order":  "overall",
	}
	err = client.RequestWithParams("GET", "/contact/", nil, params, &contacts)
	return
}

// https://developers.podio.com/doc/contacts/get-user-contact-60514
func (client *Client) GetContact(userId int64) (contact Contact, err error) {
	path := fmt.Sprintf("/contact/user/%d", userId)
	err = client.Request("GET", path, nil, nil, &contact)
	return
}
