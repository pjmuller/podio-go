package podio

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type AppField struct {
	Id         int64       `json:"field_id"`
	ExternalId string      `json:"external_id"`
	Type       string      `json:"type"`
	Label      string      `json:"label"`
	Status     string      `json:"status"`
	Config     FieldConfig `json:"config"`
}

type FieldConfig struct {
	Required     bool             `json:"required"`
	Delta        int              `json:"delta"`
	Description  string           `json:"description"`
	Label        string           `json:"label"` // for creating app fields we need to pass it on config level, later on when reading we will receive it in the main AppField struct (Podio is a bit inconsistent...)
	Settings     *json.RawMessage `json:"settings"`
	AlwaysHidden bool             `json:"hidden_create_view_edit"`
	Hidden       bool             `json:"hidden"`
}

type FieldRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type revisionResponse struct {
	Revision int `json:"revision"`
}

// https://developers.podio.com/doc/applications/add-new-app-field-22354
func (client *Client) CreateAppField(appId int64, params map[string]interface{}) (AppFieldId int64, err error) {
	path := fmt.Sprintf("/app/%d/field/", appId)
	var appField AppField
	err = client.RequestWithParams("POST", path, nil, params, &appField)
	AppFieldId = appField.Id

	return
}

// https://developers.podio.com/doc/applications/update-an-app-field-22356
func (client *Client) UpdateAppField(appId, appFieldId int64, params map[string]interface{}) (revision int, err error) {
	path := fmt.Sprintf("/app/%d/field/%d", appId, appFieldId)
	var resp revisionResponse
	err = client.RequestWithParams("PUT", path, nil, params, &resp)
	revision = resp.Revision
	return
}

// https://developers.podio.com/doc/applications/update-an-app-field-22356
func (client *Client) UpdateAppFieldRawConfig(appId, appFieldId int64, config json.RawMessage) (int, error) {
	path := fmt.Sprintf("/app/%d/field/%d", appId, appFieldId)
	var resp revisionResponse

	body := bytes.NewReader(config)
	_, _, _, err := client.request("PUT", path, nil, body, &resp)
	if err != nil {
		return 0, err
	}

	return resp.Revision, nil
}

// https://developers.podio.com/doc/items/get-field-ranges-24242866
func (client *Client) GetFieldRange(fieldID int64) (FieldRange, error) {
	path := fmt.Sprintf("/item/field/%d/range", fieldID)
	var resp FieldRange
	err := client.Request("GET", path, nil, nil, &resp)
	return resp, err
}
