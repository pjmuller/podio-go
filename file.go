package podio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type File struct {
	Id   int    `json:"file_id"`
	Name string `json:"name"`
	Link string `json:"link"`
	Size int    `json:"size"`
	Push Push   `json:"push"`

	Mimetype  string  `json:"mimetype"`
	Context   FileRef `json:"context"`
	CreatedBy FileRef `json:"created_by"`
	// AppFieldId 	int 		`json:"app_field_id"`
	CreatedOn string `json:"created_on"` // we keep this simple to save on processing power
}

type FileRef struct {
	Id   int64    `json:"id"`
	Type string   `json:"type"`
	Data FileData `json:"data"`
}

type FileData struct {
	App FileApp `json:"app"`
}

// we do not take the standard App struct
// to save a bit on parsing + memory
type FileApp struct {
	Id int64 `json:"app_id"`
}

// https://developers.podio.com/doc/files/get-files-4497983
func (client *Client) GetFiles() (files []File, err error) {
	err = client.Request("GET", "/file", nil, nil, &files)
	return
}

// https://developers.podio.com/doc/files/get-file-22451
func (client *Client) GetFile(fileId int) (file *File, err error) {
	err = client.Request("GET", fmt.Sprintf("/file/%d", fileId), nil, nil, &file)
	return
}

func (client *Client) GetFileContents(url string) ([]byte, error) {
	link := fmt.Sprintf("%s?oauth_token=%s", url, client.authToken.AccessToken)
	resp, err := http.Get(link)

	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func (client *Client) GetFileContentsToTempFile(url string) (tempFilePath string, close func(), err error) {
	// step 1: download the contents
	link := fmt.Sprintf("%s?oauth_token=%s", url, client.authToken.AccessToken)
	resp, err := http.Get(link)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	// step 2: create a tempfile + closing function
	tempFile, err := ioutil.TempFile(os.TempDir(), "podio_file")
	if err != nil {
		return
	}
	defer tempFile.Close()

	// extra function to remove the temp file once processed
	close = func() {
		os.Remove(tempFile.Name())
	}

	// step 3: write to file
	// io.Copy works in chunks of 32KB so no worries of memory overrun
	_, err = io.Copy(tempFile, resp.Body)
	tempFilePath = tempFile.Name()
	return
}

// https://developers.podio.com/doc/files/upload-file-1004361
func (client *Client) CreateFile(name string, contents []byte) (file *File, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("source", name)
	if err != nil {
		return nil, err
	}

	_, err = part.Write(contents)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("filename", name)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Content-Type": writer.FormDataContentType(),
	}

	err = client.Request("POST", "/file", headers, body, &file)
	return
}

// https://developers.podio.com/doc/files/replace-file-22450
func (client *Client) ReplaceFile(oldFileId, newFileId int) error {
	path := fmt.Sprintf("/file/%d/replace", newFileId)
	params := map[string]interface{}{
		"old_file_id": oldFileId,
	}

	return client.RequestWithParams("POST", path, nil, params, nil)
}

// https://developers.podio.com/doc/files/attach-file-22518
func (client *Client) AttachFile(fileId int, refType string, refId int64) error {
	path := fmt.Sprintf("/file/%d/attach", fileId)
	params := map[string]interface{}{
		"ref_type": refType,
		"ref_id":   refId,
	}

	return client.RequestWithParams("POST", path, nil, params, nil)
}

// https://developers.podio.com/doc/files/delete-file-22453
func (client *Client) DeleteFile(fileId int) error {
	path := fmt.Sprintf("/file/%d", fileId)
	return client.Request("DELETE", path, nil, nil, nil)
}

// https://developers.podio.com/doc/files/copy-file-89977
func (client *Client) CopyFile(fileId int) (int, error) {
	path := fmt.Sprintf("/file/%d/copy", fileId)
	rsp := &struct {
		FileId int `json:"file_id"`
	}{}
	err := client.Request("POST", path, nil, nil, rsp)
	return rsp.FileId, err
}

// https://developers.podio.com/doc/files/get-files-on-space-22471
func (client *Client) FindFilesForSpaceJson(spaceId int, params map[string]interface{}) (rawResponse *json.RawMessage, err error) {
	path := fmt.Sprintf("/file/space/%d/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &rawResponse)
	return
}

// https://developers.podio.com/doc/files/get-files-on-space-22471
func (client *Client) FindFilesForSpace(spaceId int64, params map[string]interface{}) (files []*File, err error) {
	path := fmt.Sprintf("/file/space/%d/", spaceId)
	err = client.RequestWithParams("GET", path, nil, params, &files)
	return
}
