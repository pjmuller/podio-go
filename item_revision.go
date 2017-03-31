package podio

import (
	"encoding/json"
	"fmt"
)

// https://developers.podio.com/doc/items/revert-to-revision-194362682
func (client *Client) RevertToRevision(ItemId int64, revisionId int) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/item/%d/revision/%d/revert_to", ItemId, revisionId)
	err = client.Request("POST", path, nil, nil, &rawResponse)
	return
}
