package podio

import (
	"encoding/json"
)

type Field struct {
	Id         int64   `json:"field_id"`
	ExternalId string  `json:"external_id"`
	Type       string  `json:"type"`
	Label      string  `json:"label"`
  Type       string  `json:"type"`
  Status     string  `json:"status"`

	Config		 FieldConfig     `json:"config"`
  ConfigJSON json.RawMessage `json:"config"`
}

type FieldConfig struct {
	Required      bool 					  `json:"required"`
	Delta 	      int 					  `json:"delta"`
  Description   string          `json:"description"`
	Settings      FieldSettings	  `json:"settings"`
  // SettingsJSON  json.RawMessage `json:"settings"`
}

type FieldSettings struct {
	ReturnType string `json:"return_type"` // for calculations
}
