package podio

import "fmt"

// Comment is a comment on an object in podio.
// The object to which this comment is associated is described in this Reference.
type Comment struct {
	Id         int64      `json:"comment_id"`
	ExternalId string     `json:"external_id"`
	RichValue  string     `json:"rich_value"`
	Value      string     `json:"value"`
	Ref        *Reference `json:"ref"`
	Files      []*File    `json:"files"`
	CreatedBy  ByLine     `json:"created_by"`
	CreatedVia Via        `json:"created_via"`
	CreatedOn  Time       `json:"created_on"`
	LastEditOn *string    `json:"last_edit_on"` // nil when never edited
	IsLiked    bool       `json:"is_liked"`
	LikeCount  int        `json:"like_count"`
	Embed      Embed      `json:"embed"`
}

// Comment adds a comment to a podio object. It returns a Comment (with podio ID) or an error if one occured.
//
// refType (item, task, ...) and refId identifies the podio object to which the comment is added.
// text is the actual comment value.
// Additional parameters can be set in the params map.
func (client *Client) Comment(refType string, refId int64, text string, params map[string]interface{}) (*Comment, error) {
	path := fmt.Sprintf("/comment/%s/%d/", refType, refId)
	if params == nil {
		params = map[string]interface{}{}
	}
	params["value"] = text

	comment := &Comment{}
	err := client.RequestWithParams("POST", path, nil, params, comment)
	return comment, err
}

// UpdateComment updates a comment in Podio
func (client *Client) UpdateComment(commentID int64, text string, params map[string]interface{}) error {
	path := fmt.Sprintf("/comment/%d/", commentID)
	if params == nil {
		params = map[string]interface{}{}
	}
	params["value"] = text

	err := client.RequestWithParams("PUT", path, nil, params, nil)
	return err
}

// DeleteComment deletes a comment in Podio
func (client *Client) DeleteComment(commentID int64) error {
	path := fmt.Sprintf("/comment/%d/", commentID)
	return client.Request("DELETE", path, nil, nil, nil)
}

// https://developers.podio.com/doc/comments/get-comments-on-object-22371
// GetComments retrieves the comments associated with a podio object.
//
// refType is the type of the podio object. For legal type values see
// refId is the podio id of the podio object.
func (client *Client) GetComments(refType string, refId int64) (comments []*Comment, err error) {
	path := fmt.Sprintf("/comment/%s/%d/", refType, refId)
	err = client.Request("GET", path, nil, nil, &comments)
	return
}

// https://developers.podio.com/doc/comments/get-a-comment-22345
func (client *Client) GetComment(commentId int64) (comment *Comment, err error) {
	path := fmt.Sprintf("/comment/%d/", commentId)
	err = client.Request("GET", path, nil, nil, &comment)
	return
}
