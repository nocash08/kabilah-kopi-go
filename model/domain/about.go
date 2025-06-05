package domain

type About struct {
	Id         uint   `json:"id"`
	Heading    string `json:"heading"`
	Subheading string `json:"subheading"`
}
