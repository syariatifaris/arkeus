package entity

type RenderResult struct {
	HttpContentType string `json:"http_content_type"`
	StatusCode      int    `json:"status_code"`
	Body            string `json:"body"`
}
