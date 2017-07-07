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
	Label       string           `json:"label"` // for creating app fields we need to pass it on config level, later on when reading we will receive it in the main AppField struct (Podio is a bit inconsistent...)
	Settings    *json.RawMessage `json:"settings"`
}
