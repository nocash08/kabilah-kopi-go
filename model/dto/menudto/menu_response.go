package menudto

type MenuResponse struct {
	Id         uint   `json:"id"`
	Heading    string `json:"heading"`
	Subheading string `json:"subheading"`
	Thumbnail  string `json:"thumbnail"`
}
