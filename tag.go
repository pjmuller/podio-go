package podio

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// https://developers.podio.com/doc/tags/create-tags-22464
func (client *Client) CreateTags(refType string, refId int64, tags []string) (err error) {
	path := fmt.Sprintf("/tag/%s/%d/", refType, refId)

	buf, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	body := bytes.NewReader(buf)

	err = client.Request("POST", path, nil, body, nil)
	return
}
