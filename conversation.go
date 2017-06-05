package podio

type Conversation struct {
	Id           int64                     `json:"conversation_id"`
	Subject      string                    `json:"subject"`
	CreatedOn    Time                      `json:"created_on"`
	Messages     []ConversationMessage     `json:"messages"`
	Participants []ConversationParticipant `json:"participants"`
	// Session ConversationSession `json:"session"`
}

type ConversationMessage struct {
	Id        int64     `json:"message_id"`
	Text      string    `json:"text"`
	CreatedOn Time      `json:"created_on"`
	CreatedBy RefSimple `json:"created_by"`
	Files     []File    `json:"files"`
	Embed     Embed     `json:"embed"`
	EmbedFile File      `json:"embed_file"`
}

type ConversationParticipant struct {
	Id     int64  `json:"user_id"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

// https://developers.podio.com/doc/conversations/create-conversation-v2-37301474
func (client *Client) CreateConversation(params map[string]interface{}) (c Conversation, err error) {
	err = client.RequestWithParams("POST", "/conversation/v2/", nil, params, &c)
	return
}
