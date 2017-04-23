package podio

// User contains account information
type User struct {
	Id        int    `json:"user_id"`
	Mail      string `json:"mail"`
	Status    string `json:"status"` // "inactive", "active" or "blacklisted"
	Locale    string `json:"locale"` // http://en.wikipedia.org/wiki/ISO_639-1
	Timezone  string `json:"timezone"`
	CreatedOn Time   `json:"created_on"`
}

// GetUser gets account information for current connected user
// https://developers.podio.com/doc/users/get-user-22378
func (client *Client) GetUser() (user User, err error) {
	err = client.Request("GET", "/user", nil, nil, &user)
	return
}
