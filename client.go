package podio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	httpClient *http.Client
	authToken  *AuthToken
}

type Error struct {
	Parameters interface{} `json:"error_parameters"`
	Detail     interface{} `json:"error_detail"`
	Propagate  bool        `json:"error_propagate"`
	Request    struct {
		URL   string `json:"url"`
		Query string `json:"query_string"`
	} `json:"request"`
	Description string `json:"error_description"`
	Type        string `json:"error"`
}

func (p *Error) Error() string {
	return fmt.Sprintf("%s: %s", p.Type, p.Description)
}

func NewClient(authToken *AuthToken) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
		authToken: authToken,
	}
}

func (client *Client) Request(method string, path string, headers map[string]string, body io.Reader, out interface{}) error {
	_, _, _, err := client.request(method, path, headers, body, out)
	return err
}

func (client *Client) RequestWithParams(method string, path string, headers map[string]string, params map[string]interface{}, out interface{}) error {
	_, _, _, err := client.requestWithParams(method, path, headers, params, out)
	return err
}

func (client *Client) requestWithParamsAndStatusCode(method string, path string, headers map[string]string, params map[string]interface{}, out interface{}) (int, error) {
	statusCode, _, _, err := client.requestWithParams(method, path, headers, params, out)
	return statusCode, err
}

func (client *Client) requestWithParamsAndRemainingLimit(method string, path string, headers map[string]string, params map[string]interface{}, out interface{}) (int, int, error) {
	_, remaining, limit, err := client.requestWithParams(method, path, headers, params, out)
	return remaining, limit, err
}

func (client *Client) request(method string, path string, headers map[string]string, body io.Reader, out interface{}) (int, int, int, error) {
	// for some reason `httpClient: &http.Client{Timeout: 5 * time.Minute}` doesn't seem to work, so trying with this extra line
	ctx, cncl := context.WithTimeout(context.Background(), time.Minute*5)
	defer cncl()

	req, err := http.NewRequestWithContext(ctx, method, "https://api.podio.com"+path, body)
	if err != nil {
		return 0, 0, 0, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	req.Header.Add("Authorization", "OAuth2 "+client.authToken.AccessToken)
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return 0, 0, 0, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, 0, err
	}

	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		podioErr := &Error{}
		err := json.Unmarshal(respBody, podioErr)
		if err != nil {
			return 0, 0, 0, errors.New(string(respBody))
		}
		return 0, 0, 0, podioErr
	}

	limitString := resp.Header.Get("X-Rate-Limit-Limit")
	remainingString := resp.Header.Get("X-Rate-Limit-Remaining")
	limit, _ := strconv.Atoi(limitString)
	remaining, _ := strconv.Atoi(remainingString)

	if out != nil {
		err := json.Unmarshal(respBody, out)
		return 0, remaining, limit, err
	}

	return resp.StatusCode, remaining, limit, nil
}

func (client *Client) requestWithParams(method string, path string, headers map[string]string, params map[string]interface{}, out interface{}) (int, int, int, error) {
	var body io.Reader

	if method == "GET" {
		pathURL, err := url.Parse(path)
		if err != nil {
			return 0, 0, 0, err
		}
		query := pathURL.Query()
		for key, value := range params {
			query.Add(key, fmt.Sprint(value))
		}
		pathURL.RawQuery = query.Encode()
		path = pathURL.String()
	} else {
		buf, err := json.Marshal(params)
		if err != nil {
			return 0, 0, 0, err
		}
		body = bytes.NewReader(buf)
	}

	respCode, rateLimitRemaining, rateLimit, err := client.request(method, path, headers, body, out)
	return respCode, rateLimitRemaining, rateLimit, err
}

func (client *Client) AddOptionsToPath(path string, options map[string]interface{}) (string, error) {
	pathURL, err := url.Parse(path)
	if err != nil {
		return path, err
	}
	query := pathURL.Query()
	for key, value := range options {
		query.Add(key, fmt.Sprint(value))
	}
	pathURL.RawQuery = query.Encode()
	return pathURL.String(), err
}
