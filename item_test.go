package podio

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnmarshalTextValue(t *testing.T) {
	r := require.New(t)

	fieldJson := []byte(`{
		"type": "text",
		"values": [{
			"value": "a"
		}]
	}`)

	field := &Field{}
	err := json.Unmarshal(fieldJson, field)
	r.NoError(err)

	values, ok := field.Values.([]TextValue)
	r.True(ok, "Expected values to be []podio.TextValue, is %#v", field.Values)
	r.Len(values, 1)
	r.Equal(values[0].Value, "a")
}
