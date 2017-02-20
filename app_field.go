package podio

import "encoding/json"

type AppField struct {
	Id         int64       `json:"field_id"`
	ExternalId string      `json:"external_id"`
	Type       string      `json:"type"`
	Label      string      `json:"label"`
	Status     string      `json:"status"`
	Config     FieldConfig `json:"config"`
}

type FieldConfig struct {
	Required    bool             `json:"required"`
	Delta       int              `json:"delta"`
	Description string           `json:"description"`
	Settings    *json.RawMessage `json:"settings"`
}
