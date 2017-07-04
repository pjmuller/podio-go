package podio

import "fmt"

// View defines
type View struct {
	SortBy   interface{}          `json:"sort_by"`   // app field id OR meta attributes (app_item_id, ...). Default = created_on
	SortDesc bool                 `json:"sort_desc"` // by default true
	Filters  []viewFilter         `json:"filters"`
	Fields   map[string]viewField `json:"fields"` // which columns do we show
}

type viewFilter struct {
	Key             interface{}                `json:"key"`
	Values          interface{}                `json:"values"`           // depends on key we are filtering on. cat and items is with array of IDs, dates is with {from: ... to: ...}, etc
	HumanizedValues []viewHumanizedFilterValue `json:"humanized_values"` // translate the IDs / date ranges used in Values into human readable text
}

type viewHumanizedFilterValue struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
}

// WATCH OUT: when a field was never edited (so hidden = false, width = 200)
// then it will not be included into the
//
// ALSO as values are passed as a map, we can't use the order and always need to resort based on the app_field.delta (!not the viewField.delta)
type viewField struct {
	DeltaOffset int     `json:"delta_offset"` // offset from the fields normal delta (typically 0)
	Width       int     `json:"width"`        // default 200
	Hidden      bool    `json:"hidden"`       // True if the field is hidden
	Use         *string `json:"use"`          // for card view: use of the column either "x_axis" or "y_axis". Else null
}

// Saved views can show subgroups. Useful for quick navigation
type viewGrouping struct {
	Type     string      `json:"type"`      // "field" or "revision"
	Value    interface{} `json:"value"`     //  field_id in case of "field" type, "created_by", "created_on" or "tags" in case of "revision",
	SubValue *string     `json:"sub_value"` // for date fields: "date", "weekday", "week", "month" or "year"
}

// the actual values of the grouping
type viewGroupingCounts struct {
	Total  int                 `json:"total"` // total count of items in all groups,
	Groups []viewGroupingCount `json:"groups"`
}

type viewGroupingCount struct {
	Count  int         `json:"total"`  // items count of the single group
	Avatar *File       `json:"avatar"` // user avatar file when grouping by contact or created_by, otherwise null
	Color  *string     `json:"color"`  // color of a category option when grouping by category field, otherwise null
	Value  interface{} `json:"value"`  // a unique value for each group (typically id)
	Label  string      `json:"label"`  // a text label for each group
}

// https://developers.podio.com/doc/views/get-view-27450
func (client *Client) GetView(appID int64, viewIdOrName interface{}) (v View, err error) {
	path := fmt.Sprintf("/view/app/%d/%v", appID, viewIdOrName)
	err = client.Request("GET", path, nil, nil, &v)
	return
}
