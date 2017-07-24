package podio

// Embed describes a Podio embed object
type Embed struct {
	Id          int    `json:"embed_id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EmbedHTML   string `json:"embed_html"`
	URL         string `json:"url"`
	OriginalURL string `json:"original_url"`
	ResolvedURL string `json:"resolved_url"`
	Hostname    string `json:"hostname"`
	EmbedHeight int    `json:"embed_height"`
	EmbedWidth  int    `json:"embed_width"`
}

type EmbedSimple struct {
	Id              int    `json:"podio_id"`
	Title           string `json:"title"`
	URL             string `json:"url"`
	ThumbnailFileId int    `json:"thumbnail_file_id"`
}

// https://developers.podio.com/doc/embeds/add-an-embed-726483
func (client *Client) CreateEmbed(params map[string]interface{}) (embed *Embed, err error) {
	err = client.RequestWithParams("POST", "/embed/", nil, params, &embed)
	return
}
