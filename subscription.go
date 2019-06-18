package podio

import "fmt"

// https://developers.podio.com/doc/subscriptions/unsubscribe-by-reference-22410
func (client *Client) DeleteSubscription(refType string, refId int64) error {
	path := fmt.Sprintf("/subscription/%s/%d", refType, refId)
	return client.Request("DELETE", path, nil, nil, nil)
}
