package domain

type Menu struct {
	Id         uint   `json:"id"`
	Heading    string `json:"heading"`
	Subheading string `json:"subheading"`
	Thumbnail  string `json:"thumbnail"`
}
