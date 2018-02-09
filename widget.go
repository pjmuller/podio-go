package podio

import (
	"fmt"
)

// Widget are small components that can be installed on an organization, space, app or user. Every widget has a title and is of a certain type.
type Widget struct {
	ID        int64        `json:"widget_id"`
	Ref       RefSimple    `json:"ref"`
	Type      string       `json:"type"` // text / calculation / text / image / link / tag_cloud / tasks / events / profiles / apps / app_view / contacts / files
	Title     string       `json:"title"`
	Config    WidgetConfig `json:"config"`
	CreatedOn Time         `json:"created_on"`
	CreatedBy ByLine       `json:"created_by"`
	Cols      int          `json:"cols"`
	Rows      int          `json:"rows"`
	X         int          `json:"x"`
	Y         int          `json:"y"`
	Data      WidgetData   `json:"data"`
}

// WidgetConfig holds (for now) mainly the WidgetCalculationConfig (later should be split up just like with do with item.go#Field)
type WidgetConfig struct {
	Layout      string                `json:"layout"` // table
	Calculation WidgetCalculationView `json:"calculation"`
	AppID       int64                 `json:"app_id"`
	Unit        string                `json:"unit"` // user defined, e.g. "people" / "USD" / "rainbows"
}

// WidgetCalculationView is used to scope a widget to filters + grouping
type WidgetCalculationView struct {
	Sorting     string          `json:"sorting"`   // label_asc / label_des / value_asc / value_desc
	Aggregation string          `json:"sort_desc"` // "count" / "sum"
	Limit       int             `json:"limit"`     // e.g. 15 (numer of rows to show)
	Filters     []viewFilter    `json:"filters"`
	Formula     []WidgetFormula `json:"formula"`
	Grouping    viewGrouping    `json:"grouping"`
	// Groupings []viewGrouping         `json:"groupings"` // not sure how different from viewGrouping, don't think you can have more than 1 grouping
}

// WidgetFormula is used within calculation widgets to get the desired number result (typically just the field we are summing)
type WidgetFormula struct {
	Type  string      `json:"type"`  // "field" (sum (number) field) / "operator" / "number" (for manual number)
	Value interface{} `json:"value"` // "field" => field_id, "operator" => plus/minus/multiply/divide, "number" => e.g. 5
}

// WidgetData contains the rows that matched the scoped criteria
type WidgetData struct {
	Total            float64           `json:"total"` // grand total count (can be higher than limit)
	WidgetDataGroups []WidgetDataGroup `json:"groups"`
	// Files            []File            `json:"files"` // only pre
}

// WidgetDataGroup holds data for one data row
type WidgetDataGroup struct {
	Count float64 `json:"count"` // typically ints but calcs can results in decimal values
	Value string  `json:"value"` // e.g. "2017-12-31" when grouped per date
	Label string  `json:"label"` // the translated value (often same as value)
	// Color string `json:"total"`// when using category fields to group
	// Avatar *File `json:"avatar"` // when grouping per profile/user
}

// https://developers.podio.com/doc/widgets/get-widget-22489
func (client *Client) GetWidget(widgetID int64) (w Widget, err error) {
	path := fmt.Sprintf("/widget/%d", widgetID)
	err = client.Request("GET", path, nil, nil, &w)
	return
}

// https://developers.podio.com/doc/widgets/get-widgets-22494
func (client *Client) GetWidgets(refType string, refID int64) (w []Widget, err error) {
	path := fmt.Sprintf("/widget/%s/%d", refType, refID)
	err = client.Request("GET", path, nil, nil, &w)
	return
}

// https://developers.podio.com/doc/widgets/delete-widget-22492
func (client *Client) DeleteWidget(widgetID int64) (err error) {
	path := fmt.Sprintf("/widget/%d", widgetID)
	return client.Request("DELETE", path, nil, nil, nil)
}

// https://developers.podio.com/doc/widgets/create-widget-22491
func (client *Client) CreateWidget(refType string, refID int64, params map[string]interface{}) (id int64, err error) {
	path := fmt.Sprintf("/widget/%s/%d", refType, refID)

	response := struct {
		ID int64 `json:"widget_id"`
	}{}

	err = client.RequestWithParams("POST", path, nil, params, &response)
	id = response.ID
	return id, err
}

// ----------------------------------------------------------------------------
// Section: TODO later
// func (f *WidgetConfig) UnmarshalJSON(data []byte) error {
// }

// type WidgetConfig struct {
// 	Config interface{} // will point to WidgetBasicConfig / WidgetCalculationConfig
// }

// WidgetBasicConfig is used by a lot of widgets, it only needs the limit of records to show in the tile
// type WidgetBasicConfig struct {
// 	Limit int `json:"limit"`
// }

// type WidgetCalculationConfig struct {
// 	// see current WidgetConfig
// }
