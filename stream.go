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
	CreatedVia Via        `json:"created_via"`
	CreatedBy  RefSimple  `json:"created_by"`
	CreatedOn  string     `json:"created_on"`
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
	Fields   []*ValuesFromTo `json:"fields"`
	Revision int             `json:"revision"`

	// comment related
	RichValue  string      `json:"rich_value"`
	Files      []File      `json:"files"`
	Embed      EmbedSimple `json:"embed"`
	LastEditOn *string     `json:"last_edit_on"`
	// Questions []Question `json:"questions"` // to do later
}

// revisions digging deeper
type ValuesFromTo struct {
	Field          Field           `json:"field"`
	FromValuesJSON json.RawMessage `json:"from"`
	ToValuesJSON   json.RawMessage `json:"to"`
}

type FieldSimple struct {
	Id int64 `json:"field_id"`
}

type RefSimple struct {
	Type string `json:"type"` // comment / item_revisions
	Id   int64  `json:"id"`
}

// ----------------------------------------------------------------------------
// Section: Simplified version for fetching incoming references

type StreamReference struct {
	RefId          int64                    `json:"id"`
	ActivityGroups []ActivityGroupReference `json:"activity_groups"`
}

type ActivityGroupReference struct {
	Activities []ActivityReference `json:"activities"`
}

type ActivityReference struct {
	DataRef RefSimple `json:"data_ref"`
}

// ----------------------------------------------------------------------------
// Section: API calls

// https://developers.podio.com/doc/stream/get-space-stream-v3-116373969
func (client *Client) StreamForSpaceV3Json(spaceId int64, params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/stream/space/%d/v3/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/stream/get-space-stream-v3-116373969
func (client *Client) StreamForSpaceV3(spaceId int64, params map[string]interface{}) (s []Stream, err error) {
	path := fmt.Sprintf("/stream/space/%d/v3/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &s)
	return
}

// https://developers.podio.com/doc/stream/get-application-stream-v3-100406563
func (client *Client) StreamForAppV3References(appId int64, params map[string]interface{}) (s []StreamReference, err error) {
	path := fmt.Sprintf("/stream/app/%d/v3/", appId)
	err = client.RequestWithParams("GET", path, nil, params, &s)
	return
}
