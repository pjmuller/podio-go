package podio

import "encoding/json"
import "fmt"

type Hook struct {
	Id     int64  `json:"hook_id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}

// https://developers.podio.com/doc/hooks/create-hook-215056
func (client *Client) CreateHookJson(refType string, refId int64, url string, hookType string) (rawResponse *json.RawMessage, err error) {
  path := fmt.Sprintf("/hook/%s/%d", refType, refId)
  params := map[string]interface{}{
		"url": url,
    "type": hookType,
	}
  err = client.RequestWithParams("POST", path, nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/hooks/create-hook-215056
func (client *Client) CreateHook(refType string, refId int64, url string, hookType string) (hook Hook, err error) {
  path := fmt.Sprintf("/hook/%s/%d", refType, refId)
  params := map[string]interface{}{
		"url": url,
    "type": hookType,
	}
  err = client.RequestWithParams("POST", path, nil, params, &hook)
	return
}

// https://developers.podio.com/doc/hooks/request-hook-verification-215232
func (client *Client) VerifyHook(hookId int64) error {
  path := fmt.Sprintf("/hook/%d/verify/request", hookId)
  return client.Request("POST", path, nil, nil, nil)
}

// https://developers.podio.com/doc/hooks/validate-hook-verification-215241
func (client *Client) ValidateHook(hookId int64, code string) error {
  path := fmt.Sprintf("/hook/%d/verify/validate", hookId)
  params := map[string]interface{}{
		"code": code,
	}
  return client.RequestWithParams("POST", path, nil, params, nil)
}

// https://developers.podio.com/doc/hooks/delete-hook-215291
func (client *Client) DeleteHook(hookId int64) error {
  path := fmt.Sprintf("/hook/%d", hookId)
  return client.Request("DELETE", path, nil, nil, nil)
}

// https://developers.podio.com/doc/hooks/get-hooks-215285
func (client *Client) FindHooksJson(refType string, refId int64) (rawResponse *json.RawMessage, err error) {
  path := fmt.Sprintf("/hook/%s/%d", refType, refId)
  err = client.Request("GET", path, nil, nil, &rawResponse)
	return
}

// https://developers.podio.com/doc/hooks/get-hooks-215285
func (client *Client) FindHooks(refType string, refId int64) (hooks []Hook, err error) {
  path := fmt.Sprintf("/hook/%s/%d", refType, refId)
  err = client.Request("GET", path, nil, nil, &hooks)
	return
}
