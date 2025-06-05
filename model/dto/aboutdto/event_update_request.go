package aboutdto

type AboutUpdateRequest struct {
	Id         uint   `json:"id" validate:"required"`
	Heading    string `json:"heading" validate:"required,max=200,min=1"`
	Subheading string `json:"subheading" validate:"required,max=1000,min=1"`
}
