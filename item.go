package podio

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// optional custom types
// - AppValue -> reference new 'ItemRef' that has Id + existing AppSimple

// Item describes a Podio item object
type Item struct {
	Id                 int64    `json:"item_id"`
	AppItemId          int      `json:"app_item_id"`
	FormattedAppItemId string   `json:"app_item_id_formatted"`
	Title              string   `json:"title"`
	Files              []*File  `json:"files"`
	Fields             []*Field `json:"fields"`
	Space              Space    `json:"space"`
	App                App      `json:"app"`
	CreatedVia         Via      `json:"created_via"`
	CreatedBy          ByLine   `json:"by_line"`
	CreatedOn          Time     `json:"created_on"`
	Link               string   `json:"link"`
	Revision           int      `json:"revision"`
	Push               Push     `json:"push"`
}

type ItemSimple struct {
	Id           int64    `json:"item_id"`
	AppItemId    int      `json:"app_item_id"`
	Title        string   `json:"title"`
	Revision     int      `json:"revision"`
	Tags         []string `json:"tags"` // => don't think this comes along
	ExternalId   string   `json:"external_id"`
	CommentCount int      `json:"comment_count"`

	// App                AppSimple 		`json:"app"` => When filtering on app we don't get the app passed
	CreatedVia      Via          `json:"created_via"`
	CreatedBy       ByLineSimple `json:"created_by"`
	CreatedOn       Time         `json:"created_on"` // not the standard time  RFC 3339 format so can't use time.Time
	CurrentRevision RevisionInfo `json:"current_revision"`
	LastEditOn      Time         `json:"last_edit_on"`
	LastActivityOn  Time         `json:"last_event_on"`

	// values
	Fields []*Field `json:"fields"`

	// Files
	Files []*File `json:"files"`
}

type ItemMicro struct {
	Id         int64  `json:"item_id"`
	AppItemId  int    `json:"app_item_id"`
	Title      string `json:"title"`
	Revision   int    `json:"revision"`
	ExternalID string `json:"external_id"` // typically not included but we fetch it through items.view(micro).fields(external_id)
}

type ItemMini struct {
	Id        int64  `json:"item_id"`
	AppItemId int    `json:"app_item_id"`
	Title     string `json:"title"`
	Revision  int    `json:"revision"`

	// InitialRevision RevisionInfo `json:"initial_revision"`
	CreatedVia Via          `json:"created_via"`
	CreatedBy  ByLineSimple `json:"created_by"`
	CreatedOn  Time         `json:"created_on"` // not the standard time  RFC 3339 format so can't use time.Time
}

type ItemReferences struct {
	App   ItemReferenceApp   `json:"app"`
	Field ItemReferenceField `json:"field"`
	Items []*ItemReference   `json:"items"`
}

type ItemReference struct {
	ItemID int64 `json:"item_id"`
}

type ItemReferenceApp struct {
	AppID int64 `json:"app_id"`
}

type ItemReferenceField struct {
	FieldID int64 `json:"field_id"`
}

// trick to get the "LastEditOn"
type RevisionInfo struct {
	LastEditOn string `json:"created_on"`
}

type AppSimple struct {
	Id int64 `json:"app_id"`
}

type ItemCount struct {
	Count int `json:"count"`
}

type itemId struct {
	Id int64 `json:"item_id"`
}

// partialField is used for JSON unmarshalling
// it is different from AppField because
// 1) we have the json values
// 2) when using AppField in elsa we want to keep Config > Settings as raw json
//    as we can't parse for all different app fields
type PartialField struct {
	Id         int64             `json:"field_id"`
	ExternalId string            `json:"external_id"`
	Type       string            `json:"type"`
	Label      string            `json:"label"`
	ValuesJSON json.RawMessage   `json:"values"`
	Config     FieldConfigSimple `json:"config"`
}

type FieldConfigSimple struct {
	Settings FieldSettingsSimple `json:"settings"`
}

type FieldSettingsSimple struct {
	ReturnType string `json:"return_type"` // for calculations
}

// Field describes a Podio field object
type Field struct {
	PartialField
	Values interface{}
}

func (f *Field) unmarshalValuesInto(out interface{}) error {
	if err := json.Unmarshal(f.ValuesJSON, &out); err != nil {
		return fmt.Errorf("[ERR] Cannot unmarshal %s into %s: %v\n", f.ValuesJSON, reflect.TypeOf(out), err)
	}
	return nil
}

func (f *Field) UnmarshalJSON(data []byte) error {
	f.PartialField = PartialField{}
	if err := json.Unmarshal(data, &f.PartialField); err != nil {
		return err
	}

	f.UnmarshalValues()

	return nil
}

// UnmarshalValues transforms a json.RawMessage message into actual podio types (App, Date, ...)
func (f *Field) UnmarshalValues() {
	switch f.Type {
	case "app":
		values := []AppValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "date":
		values := []DateValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "text":
		values := []TextValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "tag":
		values := []TagValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "number":
		values := []NumberValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "image":
		values := []ImageValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "member":
		values := []MemberValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "contact":
		values := []ContactValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "money":
		values := []MoneyValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "progress":
		values := []ProgressValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "location":
		values := []LocationValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "video":
		values := []VideoValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "duration":
		values := []DurationValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "embed":
		values := []EmbedValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "question":
		values := []QuestionValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "category":
		values := []CategoryValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "tel":
		values := []TelValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "phone":
		values := []PhoneValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "email":
		values := []EmailValue{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	case "calculation":
		switch f.Config.Settings.ReturnType {
		case "text":
			values := []TextValue{}
			f.unmarshalValuesInto(&values)
			f.Values = values

		case "number":
			values := []NumberValue{}
			f.unmarshalValuesInto(&values)
			f.Values = values

		case "date":
			values := []DateValue{}
			f.unmarshalValuesInto(&values)
			f.Values = values
		}

	default:
		// Unknown field type
		fmt.Println("error=unknown_app_field context=podio_item level=notice type='", f.Type, "' field='", f, "'")
		values := []interface{}{}
		f.unmarshalValuesInto(&values)
		f.Values = values
	}

	f.ValuesJSON = nil
}

// TextValue is the value for fields of type `text`
type TextValue struct {
	Value string `json:"value"`
}

// TagValue is the value for fields of type `tag`
type TagValue struct {
	Value string `json:"value"`
}

// NumberValue is the value for fields of type `number`
type NumberValue struct {
	Value float64 `json:"value,string"`
}

// Image is the value for fields of type `image`
type ImageValue struct {
	Value File `json:"value"`
}

type ImageAndItem struct {
	File       File
	ItemId     int64
	AppFieldId int64
}

type ImageValueSimple struct {
	FileId int `json:"file_id"`
}

// DateValue is the value for fields of type `date`
type DateValue struct {
	Start    *Time   `json:"start"`
	End      *Time   `json:"end"`
	StartUTC *string `json:"start_utc"`
	EndUTC   *string `json:"end_utc"`
}

type DateValueSimple struct {
	Start *string `json:"start_utc,omitempty"`
	End   *string `json:"end_utc,omitempty"`
}

type DateTimeValueSimple struct {
	Start time.Time `json:"start_utc,omitempty" bson:"start_utc,omitempty"`
	End   time.Time `json:"end_utc,omitempty" bson:"end_utc,omitempty"`
}

// AppValue is the value for fields of type `app`
type AppValue struct {
	Value Item `json:"value"`
}

type AppValueSimple struct {
	ItemId int64 `json:"item_id"`
	AppId  int64 `json:"app_id"`
}

// MemberValue is the value for fields of type `member`
type MemberValue struct {
	Value int `json:"value"`
}

// ContactValue is the value for fields of type `contact`
type ContactValue struct {
	Value Contact `json:"value"`
}

// MoneyValue is the value for fields of type `money`
type MoneyValue struct {
	Value    float64 `json:"value,string"`
	Currency string  `json:"currency"`
}

// without string conversion
type MoneyValueFloat struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}

// ProgressValue is the value for fields of type `progress`
type ProgressValue struct {
	Value int `json:"value"`
}

// LocationValue is the value for fields of type `location`
type LocationValue struct {
	Value        string  `json:"value"`
	Formatted    string  `json:"formatted"`
	StreetNumber string  `json:"street_number"`
	StreetName   string  `json:"street_name"`
	PostalCode   string  `json:"postal_code"`
	City         string  `json:"city"`
	State        string  `json:"state"`
	Country      string  `json:"country"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
}

// VideoValue is the value for fields of type `video`
type VideoValue struct {
	Value int `json:"value"`
}

// DurationValue is the value for fields of type `duration`
type DurationValue struct {
	Value int `json:"value"`
}

// EmbedValue is the value for fields of type `embed`
type EmbedValue struct {
	Embed Embed `json:"embed"`
	File  File  `json:"file"`
}

// CategoryValue is the value for fields of type `category`
type CategoryValue struct {
	Value struct {
		Status string `json:"status"`
		Text   string `json:"text"`
		Id     int    `json:"id"`
		Color  string `json:"color"`
	} `json:"value"`
}

// QuestionValue is the value for fields of type `question`
type QuestionValue struct {
	Value int `json:"value"`
}

// TelValue is the value for fields of type `tel`
type TelValue struct {
	Value int    `json:"value"`
	URI   string `json:"uri"`
}

type PhoneValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type EmailValue struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// CalcationValue is the value for fields of type `calculation` (currently untyped)
type CalculationValue map[string]interface{}

type ItemList struct {
	Filtered int     `json:"filtered"`
	Total    int     `json:"total"`
	Items    []*Item `json:"items"`
}

type ItemListSimple struct {
	Filtered int           `json:"filtered"`
	Total    int           `json:"total"`
	Items    []*ItemSimple `json:"items"`
}

type ItemListMicro struct {
	Filtered int          `json:"filtered"`
	Total    int          `json:"total"`
	Items    []*ItemMicro `json:"items"`
}

type ItemListMini struct {
	Filtered int         `json:"filtered"`
	Total    int         `json:"total"`
	Items    []*ItemMini `json:"items"`
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) GetItems(appId int64) (items *ItemList, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.fields(files)", appId)
	err = client.Request("POST", path, nil, nil, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) GetItemsSimple(appId int64) (items *ItemListSimple, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.fields(files)", appId)
	err = client.Request("POST", path, nil, nil, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItems(appId int64, params map[string]interface{}) (items *ItemList, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.fields(files)", appId)
	err = client.RequestWithParams("POST", path, nil, params, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItemsSimple(appId int64, params map[string]interface{}) (items *ItemListSimple, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.fields(files)", appId)
	err = client.RequestWithParams("POST", path, nil, params, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItemsMicro(appId int64, params map[string]interface{}) (items *ItemListMicro, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.view(micro).fields(external_id)", appId)
	err = client.RequestWithParams("POST", path, nil, params, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItemsMicroWithRateLimitStats(appId int64, params map[string]interface{}) (items *ItemListMicro, rateLimitRemaining, rateLimit int, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.view(micro).fields(external_id)", appId)
	rateLimitRemaining, rateLimit, err = client.requestWithParamsAndRemainingLimit("POST", path, nil, params, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItemsMini(appId int64, params map[string]interface{}) (items *ItemListMini, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.view(mini)", appId)
	err = client.RequestWithParams("POST", path, nil, params, &items)
	return
}

// https://developers.podio.com/doc/items/filter-items-4496747
func (client *Client) FilterItemsJson(appId int64, params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/item/app/%d/filter?fields=items.fields(files)", appId)
	err = client.RequestWithParams("POST", path, nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/items/export-items-4235696
func (client *Client) ExportItems(appId int64, exportFormat string, params map[string]interface{}) (int64, error) {
	path := fmt.Sprintf("/item/app/%d/export/%s", appId, exportFormat)
	rsp := &struct {
		BatchId int64 `json:"batch_id"`
	}{}

	err := client.RequestWithParams("POST", path, nil, params, rsp)

	return rsp.BatchId, err
}

// https://developers.podio.com/doc/items/get-item-by-app-item-id-66506688
func (client *Client) GetItemByAppItemId(appId int64, formattedAppItemId string) (item *Item, err error) {
	path := fmt.Sprintf("/app/%d/item/%s", appId, formattedAppItemId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// https://developers.podio.com/doc/items/get-item-by-app-item-id-66506688
func (client *Client) GetItemSimpleByAppItemId(appId int64, formattedAppItemId string) (item *ItemSimple, err error) {
	path := fmt.Sprintf("/app/%d/item/%s", appId, formattedAppItemId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// https://developers.podio.com/doc/items/get-item-by-external-id-19556702
func (client *Client) GetItemByExternalID(appId int64, externalId string) (item *Item, err error) {
	path := fmt.Sprintf("/item/app/%d/external_id/%s", appId, externalId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// https://developers.podio.com/doc/items/get-item-22360
func (client *Client) GetItem(itemId int64) (item *Item, err error) {
	path := fmt.Sprintf("/item/%d?fields=files", itemId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// get item (and more specifically app fields) in the format Elsa Understands
// https://developers.podio.com/doc/items/get-item-22360
func (client *Client) GetItemSimple(itemId int64) (item *ItemSimple, err error) {
	path := fmt.Sprintf("/item/%d", itemId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// get item with only micro attributes (FYI: there is no way to get a trimmed version from the API, but at least we don't parse all the values)
// https://developers.podio.com/doc/items/get-item-22360
func (client *Client) GetItemMicro(itemId int64) (item *ItemMicro, err error) {
	path := fmt.Sprintf("/item/%d", itemId)
	err = client.Request("GET", path, nil, nil, &item)
	return
}

// get Item with only basic info -> failed expirement, always returns all info
// https://developers.podio.com/doc/items/get-item-22360
// func (client *Client) GetItemMicro(itemId int64) (item *ItemMicro, err error) {
// func (client *Client) GetItemMicro(itemId int64) (rawResponse *json.RawMessage, err error) {
// 	path := fmt.Sprintf("/item/%d?fields=items.view(micro).fields(external_id)", itemId) // ?view=micro / fields=app.view(micro)"
// 	// 	err = client.Request("GET", path, nil, nil, &item)
// 	err = client.Request("GET", path, nil, nil, &rawResponse)
// 	return
// }

// https://developers.podio.com/doc/items/add-new-item-22362
func (client *Client) CreateItem(appId int, externalId string, fieldValues map[string]interface{}) (int64, error) {
	path := fmt.Sprintf("/item/app/%d", appId)
	params := map[string]interface{}{
		"fields": fieldValues,
	}

	if externalId != "" {
		params["external_id"] = externalId
	}

	rsp := &struct {
		ItemId int64 `json:"item_id"`
	}{}
	err := client.RequestWithParams("POST", path, nil, params, rsp)

	return rsp.ItemId, err
}

// https://developers.podio.com/doc/items/add-new-item-22362
func (client *Client) CreateItemThroughParams(appId int64, params map[string]interface{}, options map[string]interface{}) (item *ItemSimple, err error) {
	path := fmt.Sprintf("/item/app/%d", appId)
	path, err = client.AddOptionsToPath(path, options)
	err = client.RequestWithParams("POST", path, nil, params, &item)
	return
}

// https://developers.podio.com/doc/items/add-new-item-22362
func (client *Client) CreateItemJson(appId int, params map[string]interface{}, options map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/item/app/%d", appId)
	path, err = client.AddOptionsToPath(path, options)
	err = client.RequestWithParams("POST", path, nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/items/update-item-22363
func (client *Client) UpdateItem(itemId int, fieldValues map[string]interface{}) error {
	path := fmt.Sprintf("/item/%d", itemId)
	params := map[string]interface{}{
		"fields": fieldValues,
	}

	return client.RequestWithParams("PUT", path, nil, params, nil)
}

// https://developers.podio.com/doc/items/update-item-22363
func (client *Client) UpdateItemWithParams(itemId int64, params map[string]interface{}, options map[string]interface{}) (err error) {
	path := fmt.Sprintf("/item/%d", itemId)
	path, err = client.AddOptionsToPath(path, options)
	err = client.RequestWithParams("PUT", path, nil, params, nil)
	return
}

// https://developers.podio.com/doc/items/update-item-22363
func (client *Client) UpdateItemWithParamsAndStatusCode(itemId int64, params map[string]interface{}, options map[string]interface{}) (statusCode int, err error) {
	path := fmt.Sprintf("/item/%d", itemId)
	path, err = client.AddOptionsToPath(path, options)
	statusCode, err = client.requestWithParamsAndStatusCode("PUT", path, nil, params, nil)
	return
}

// https://developers.podio.com/doc/items/update-item-22363
func (client *Client) UpdateItemWithParamsAndRemainingRateLimit(itemId int64, params map[string]interface{}, options map[string]interface{}) (rateLimitRemaining, rateLimit int, err error) {
	path := fmt.Sprintf("/item/%d", itemId)
	path, err = client.AddOptionsToPath(path, options)
	rateLimitRemaining, rateLimit, err = client.requestWithParamsAndRemainingLimit("PUT", path, nil, params, nil)
	return
}

// https://developers.podio.com/doc/items/get-item-count-34819997
func (client *Client) ItemCount(appId int64, options map[string]interface{}) (count ItemCount, err error) {
	path := fmt.Sprintf("/item/app/%d/count", appId)
	path, err = client.AddOptionsToPath(path, options)

	err = client.Request("GET", path, nil, nil, &count)
	return
}

// https://developers.podio.com/doc/items/find-referenceable-items-22485
func (client *Client) ItemSearchField(AppFieldId int64, options map[string]interface{}) (items []Item, err error) {
	path := fmt.Sprintf("/item/field/%d/find", AppFieldId)
	path, err = client.AddOptionsToPath(path, options)

	err = client.Request("GET", path, nil, nil, &items)
	return
}

// https://developers.podio.com/doc/items/clone-item-37722742
func (client *Client) ItemClone(itemID int64, options map[string]interface{}) (clonedItemID itemId, err error) {
	path := fmt.Sprintf("/item/%d/clone", itemID)
	err = client.RequestWithParams("POST", path, nil, options, &clonedItemID)
	return
}

// https://developers.podio.com/doc/items/bulk-delete-items-19406111
// todo later parse the response (deleted / pending item ids)
func (client *Client) ItemBulkDelete(appID int64, params map[string]interface{}) (err error) {
	path := fmt.Sprintf("/item/app/%d/delete", appID)
	err = client.RequestWithParams("POST", path, nil, params, nil)
	return
}

// https://developers.podio.com/doc/items/delete-item-22364
func (client *Client) ItemDelete(itemID int64, params map[string]interface{}) (err error) {
	path := fmt.Sprintf("/item/%d", itemID)
	err = client.RequestWithParams("DELETE", path, nil, params, nil)
	return
}

// https://developers.podio.com/doc/items/get-item-references-22439
func (client *Client) GetItemReferences(itemID int64) (references []*ItemReferences, err error) {
	path := fmt.Sprintf("/item/%d/reference", itemID)

	err = client.Request("GET", path, nil, nil, &references)
	return
}
