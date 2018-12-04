package podio

import "fmt"

type Space struct {
	Id   int64  `json:"space_id"`
	Slug string `json:"url_label"`
	Name string `json:"name"`
	URL  string `json:"url"`
	// URLLabel string `json:"url_label"`
	OrgId    int64  `json:"org_id"`
	Push     Push   `json:"push"`
	Role     string `json:"role"`
	Archived bool   `json:"archived"`
}

type spaceIdResponse struct {
	Id  int64  `json:"space_id"`
	Url string `json:"url"`
}

func (client *Client) GetSpaces(orgId int64) (spaces []Space, err error) {
	path := fmt.Sprintf("/org/%d/space", orgId)
	err = client.Request("GET", path, nil, nil, &spaces)
	return
}

func (client *Client) GetSpace(id int64) (space *Space, err error) {
	path := fmt.Sprintf("/space/%d", id)
	err = client.Request("GET", path, nil, nil, &space)
	return
}

func (client *Client) GetSpaceByOrgIdAndSlug(orgId int64, slug string) (space *Space, err error) {
	path := fmt.Sprintf("/space/org/%d/%s", orgId, slug)
	err = client.Request("GET", path, nil, nil, &space)
	return
}

// https://developers.podio.com/doc/spaces/create-space-22390
func (client *Client) CreateSpace(orgId int64, name string) (spaceId int64, spaceUrl string, err error) {
	params := map[string]interface{}{"org_id": orgId, "name": name, "privacy": "closed", "auto_join": false, "post_on_new_app": true, "post_on_new_member": true}
	var resp spaceIdResponse
	err = client.RequestWithParams("POST", "/space/", nil, params, &resp)
	spaceId = resp.Id
	spaceUrl = resp.Url

	return
}

// https://developers.podio.com/doc/spaces/update-space-22391
func (client *Client) UpdateSpace(spaceId int64, name string) (err error) {
	path := fmt.Sprintf("/space/%d", spaceId)
	params := map[string]interface{}{"name": name}
	err = client.RequestWithParams("PUT", path, nil, params, nil)

	return err
}

// https://developers.podio.com/doc/spaces/update-space-22391
func (client *Client) UpdateSpaceUrlLabel(spaceId int64, urlLabel string) (err error) {
	path := fmt.Sprintf("/space/%d", spaceId)
	params := map[string]interface{}{"url_label": urlLabel}
	err = client.RequestWithParams("PUT", path, nil, params, nil)

	return err
}
