package podio

import "fmt"

type BatchIDResp struct {
	Id int64 `json:"batch_id"`
}

// https://developers.podio.com/doc/importer/import-app-items-212899
func (client *Client) Importer(appId int64, fileId int, params map[string]interface{}) (batchID int64, err error) {
	path := fmt.Sprintf("/importer/%d/item/app/%d", fileId, appId)
	var r BatchIDResp
	err = client.RequestWithParams("POST", path, nil, params, &r)
	batchID = r.Id
	return
}
