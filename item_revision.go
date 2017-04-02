package podio

import (
	"encoding/json"
	"fmt"
)

type ItemRevision struct {
	Id         int64     `json:"item_revision_id"`
	Revision   int       `json:"revision"`
	CreatedVia Via       `json:"created_via"`
	CreatedOn  Time      `json:"created_on"` // not the standard time  RFC 3339 format so can't use time.Time
	CreatedBy  RefSimple `json:"created_by"`
	Type       string    `json:"type"`
}

// https://developers.podio.com/doc/items/revert-to-revision-194362682
func (client *Client) RevertToRevision(ItemId int64, revisionId int) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/item/%d/revision/%d/revert_to", ItemId, revisionId)
	err = client.Request("POST", path, nil, nil, &rawResponse)
	return
}

// https://developers.podio.com/doc/items/get-item-revisions-22372
func (client *Client) RevisionsByItemId(ItemId int64) (revisions []ItemRevision, err error) {
	path := fmt.Sprintf("/item/%d/revision/", ItemId)
	err = client.Request("GET", path, nil, nil, &revisions)
	return
}
