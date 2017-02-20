package podio

import (
	"encoding/json"
	"fmt"
)

type Stream struct {
	RefId          int64           `json:"id"`
	RefType        string          `json:"type"`
	CreatedBy      RefSimple       `json:"created_by"`
	ActivityGroups []ActivityGroup `json:"activity_groups"`
	CreatedOn      string          `json:"created_on"`
	App            StreamApp       `json:"app"`
}

type StreamApp struct {
	Id int64 `json:"app_id"`
}

type ActivityGroup struct {
	Activities []Activity `json:"activities"`
}

type Activity struct {
	CreatedOn string    `json:"created_on"`
	Type      string    `json:"type"` // 'update' / 'comment'
	Data      ActData   `json:"data"`
	DataRef   RefSimple `json:"data_ref"`
}

// we are interested in both comment as revision activities
// and will keep values for both
type ActData struct {
	// revision related
	Fields   []FieldFromTo `json:"fields"`
	Revision int           `json:"revision"`

	// comment related
	RichValue  string      `json:"rich_value"`
	Files      []File      `json:"files"`
	Embed      EmbedSimple `json:"embed"`
	LastEditOn *string     `json:"last_edit_on"`
	// Questions []Question `json:"questions"` // to do later
}

// revisions digging deeper
type FieldFromTo struct {
	Field FieldSimple     `json:"field"`
	From  json.RawMessage `json:"from"`
	To    json.RawMessage `json:"to"`
}

type FieldSimple struct {
	Id int64 `json:"field_id"`
}

type RefSimple struct {
	Type string `json:"type"` // comment / item_revisions
	Id   int64  `json:"id"`
}

// https://developers.podio.com/doc/stream/get-space-stream-v3-116373969
func (client *Client) StreamForSpaceV3Json(spaceId int64, params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/stream/space/%d/v3/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &rawResponse)
	return
}

// hhttps://developers.podio.com/doc/stream/get-space-stream-v3-116373969
func (client *Client) StreamForSpaceV3(spaceId int64, params map[string]interface{}) (s []Stream, err error) {
	path := fmt.Sprintf("/stream/space/%d/v3/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &s)
	return
}
