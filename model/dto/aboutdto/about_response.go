package aboutdto

type AboutResponse struct {
	Id         uint   `json:"id"`
	Heading    string `json:"heading"`
	Subheading string `json:"subheading"`
}
