package podio

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Tag represents both label as count occurence
type Tag struct {
	Text  string `json:"text"`
	Count int    `json:"count"`
}

type TaggedObject struct {
	Type string `json:"type"`
	Id   int64  `json:"id"`
	// Title string  `json:"title"`
	// Link      string `json:"link"`
	// CreatedOn Time   `json:"created_on"`
}

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

// https://developers.podio.com/doc/tags/update-tags-39859
func (client *Client) UpdateTags(refType string, refId int64, tags []string) (err error) {
	path := fmt.Sprintf("/tag/%s/%d/", refType, refId)

	buf, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	body := bytes.NewReader(buf)

	err = client.Request("PUT", path, nil, body, nil)
	return
}

// https://developers.podio.com/doc/tags/get-tags-on-app-top-68485
func (client *Client) ListTopTagsForApp2(appId int64, query string, limit int) (tags []string, err error) {
	path := fmt.Sprintf("/tag/app/%d/top/?limit=%d&text=%s", appId, limit, query)
	err = client.RequestWithParams("GET", path, nil, nil, &tags)
	return
}

// https://developers.podio.com/doc/tags/get-tags-on-app-22467
func (client *Client) ListTagsForApp(appId int64, query string, limit int) (tags []*Tag, err error) {
	path := fmt.Sprintf("/tag/app/%d/?limit=%d&text=%s", appId, limit, query)
	err = client.RequestWithParams("GET", path, nil, nil, &tags)
	return
}

// https://developers.podio.com/doc/tags/get-objects-on-app-with-tag-22469
func (client *Client) ObjectsOnAppWithTag(appId int64, tag string) (tags []*TaggedObject, err error) {
	path := fmt.Sprintf("/tag/app/%d/search/?text=%s", appId, tag)
	err = client.RequestWithParams("GET", path, nil, nil, &tags)
	return
}

// https://developers.podio.com/doc/tags/remove-tag-22465
func (client *Client) DeleteTag(refType string, refId int64, text string) (err error) {
	// path := fmt.Sprintf("/tag/%s/%d/?text=%s", refType, refId, text)
	path := fmt.Sprintf("/tag/%s/%d/", refType, refId)
	options := map[string]interface{}{"text": text}
	path, err = client.AddOptionsToPath(path, options)
	err = client.Request("DELETE", path, nil, nil, nil)

	return
}
