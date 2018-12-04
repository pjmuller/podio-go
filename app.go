package podio

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Id            int64  `json:"app_id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	DefaultViewId int    `json:"default_view_id"`
	URLAdd        string `json:"url_add"`
	// IconId          int    `json:"icon_id"`
	LinkAdd         string `json:"link_add"`
	CurrentRevision int    `json:"current_revision"`
	// ItemName        string `json:"item_name"`
	Link     string `json:"link"`
	URL      string `json:"url"`
	URLLabel string `json:"url_label"`
	SpaceId  int    `json:"space_id"`
	Icon     string `json:"icon"`
	IconId   int    `json:"icon_id"`
	APIToken string `json:"token"`

	Fields  []AppField `json:"fields"`
	Config  AppConfig  `json:"config"`
	Layouts AppLayouts `json:"layouts"`
	Owner   AppRef     `json:"owner"`
}

type AppConfig struct {
	Name        string  `json:"name"`
	ItemName    string  `json:"item_name"`
	Description string  `json:"description"`
	Usage       *string `json:"usage"`
	Type        string  `json:"type"`         // standard
	DefaultView string  `json:"default_view"` // badge / list
	Icon        string  `json:"icon"`         // "141.png"
	IconID      int     `json:"icon_id"`      // 141
	ExternalID  *string `json:"external_id"`

	AllowEdit     bool `json:"allow_edit"`
	AllowCreate   bool `json:"allow_create"`
	SilentCreates bool `json:"silent_creates"`
	SilentEdits   bool `json:"silent_edits"`

	ShowAppItemId    bool    `json:"show_app_item_id"`
	AppItemIdPrefex  *string `json:"app_item_id_prefix"`
	AppItemIdPadding int     `json:"app_item_id_padding"`

	AllowTags            bool `json:"allow_tags"`
	AllowComments        bool `json:"allow_comments"`
	AllowAttachments     bool `json:"allow_attachments"`
	DisableNotifications bool `json:"disable_notifications"`

	RSVP           bool    `json:"rsvp"`
	RSVPLabel      *string `json:"rsvp_label"`
	YesNo          bool    `json:"yesno"`
	YesNoLabel     *string `json:"yesno_label"`
	Thumbs         bool    `json:"thumbs"`
	ThumbsLabel    *string `json:"thumbs_label"`
	Fivestar       bool    `json:"fivestar"`
	FivestartLabel *string `json:"fivestar_label"`

	Approved                bool `json:"approved"`
	CalendarColorCategoryId *int `json:"calendar_color_category_field_id"`
	// attributes not done yet: tasks
}

type AppLayouts struct {
	Badge        AppLayout `json:"badge"`
	Relationship AppLayout `json:"relationship"`
}

type AppLayout struct {
	Fields json.RawMessage `json:"fields"`
}

type AppRef struct {
	Id   int64  `json:"id"`
	Type string `json:"type"`
}

// when we create an app we only get the id
type appIdResponse struct {
	Id int64 `json:"app_id"`
}

// https://developers.podio.com/doc/applications/get-apps-by-space-22478
func (client *Client) GetApps(spaceId int64, options map[string]interface{}) (apps []*App, err error) {
	path := fmt.Sprintf("/app/space/%d", spaceId)
	err = client.RequestWithParams("GET", path, nil, options, &apps)
	return
}

// https://developers.podio.com/doc/applications/get-apps-by-space-22478
func (client *Client) GetAppsJson(spaceId int64, options map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/app/space/%d", spaceId)
	err = client.RequestWithParams("GET", path, nil, options, &rawResponse)
	return
}

// https://developers.podio.com/doc/applications/get-app-22349
func (client *Client) GetApp(id int64) (app *App, err error) {
	path := fmt.Sprintf("/app/%d", id)
	err = client.Request("GET", path, nil, nil, &app)
	return
}

// https://developers.podio.com/doc/applications/get-app-on-space-by-url-label-477105
func (client *Client) GetAppBySpaceIdAndSlug(spaceId int64, slug string) (app *App, err error) {
	path := fmt.Sprintf("/app/space/%d/%s", spaceId, slug)
	err = client.Request("GET", path, nil, nil, &app)
	return
}

// https://developers.podio.com/doc/applications/get-space-app-dependencies-45779
func (client *Client) GetSpaceDependencies(spaceId int64) (response *interface{}, err error) {
	path := fmt.Sprintf("/space/%d/dependencies", spaceId)
	err = client.Request("GET", path, nil, nil, &response)
	return
}

// https://developers.podio.com/doc/applications/add-new-app-22351
func (client *Client) CreateApp(spaceId int64, config map[string]interface{}, fields []AppField) (AppId int64, err error) {
	params := map[string]interface{}{"space_id": spaceId, "config": config, "fields": fields}
	var resp appIdResponse
	err = client.RequestWithParams("POST", "/app/", nil, params, &resp)
	AppId = resp.Id

	return
}

// https://developers.podio.com/doc/applications/update-app-22352
func (client *Client) UpdateApp(appId int64, config map[string]interface{}) (err error) {
	path := fmt.Sprintf("/app/%d", appId)
	params := map[string]interface{}{"config": config}
	err = client.RequestWithParams("PUT", path, nil, params, nil)

	return err
}

// https://developers.podio.com/doc/applications/update-app-22352
func (client *Client) UpdateAppRaw(appId int64, configRaw json.RawMessage) (err error) {
	path := fmt.Sprintf("/app/%d", appId)
	params := map[string]interface{}{"config": configRaw}
	err = client.RequestWithParams("PUT", path, nil, params, nil)

	return err
}

// https://developers.podio.com/doc/applications/install-app-22506
func (client *Client) InstallApp(appId, spaceId int64, features []string) (AppId int64, err error) {
	// when features is empty, will default to filters ['widgets', 'integration', 'forms', 'flows', 'votings'] (so all except 'items')
	path := fmt.Sprintf("/app/%d/install", appId)
	params := map[string]interface{}{"space_id": spaceId, "features": features}

	var resp appIdResponse
	err = client.RequestWithParams("POST", path, nil, params, &resp)
	AppId = resp.Id

	return
}
