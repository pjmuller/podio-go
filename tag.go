package podio

import "fmt"

// https://developers.podio.com/doc/tags/create-tags-22464
func (client *Client) CreateTags(refType string, refId int64, tags []string) (err error) {
	path := fmt.Sprintf("/tag/%s/%d/", refType, refId)
	params := make(map[string]interface{})
	for _, tag := range tags {
		params[tag] = tag
	}
	err = client.RequestWithParams("POST", path, nil, params, nil)
	return
}
