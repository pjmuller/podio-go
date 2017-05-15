package podio

import "fmt"

type Organization struct {
	Id        int64    `json:"org_id"`
	Slug      string   `json:"url_label"`
	Name      string   `json:"name"`
	Image     OrgImage `json:"image"`
	Spaces    []Space  `json:"spaces"`
	Rank      int      `json:"rank"`
	Role      string   `json:"role"`
	UserLimit int      `json:"user_limit"`
	Tier      string   `json:"tier"`
	Domains   []string `json:"domains"`
}

type OrgImage struct {
	ThumbnailLink string `json:"thumbnail_link"`
}

// https://developers.podio.com/doc/organizations/get-organizations-22344
func (client *Client) GetOrganizations() (orgs []Organization, err error) {
	err = client.Request("GET", "/org", nil, nil, &orgs)
	return
}

// https://developers.podio.com/doc/organizations/get-organization-22383
func (client *Client) GetOrganization(id int64) (org *Organization, err error) {
	path := fmt.Sprintf("/org/%d", id)
	err = client.Request("GET", path, nil, nil, &org)
	return
}

// https://developers.podio.com/doc/organizations/get-organization-by-url-22384
func (client *Client) GetOrganizationBySlug(slug string) (org *Organization, err error) {
	path := fmt.Sprintf("/org/url?org_slug=%s", slug)
	err = client.Request("GET", path, nil, nil, &org)
	return
}
