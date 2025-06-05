package menudto

import "mime/multipart"

type MenuCreateRequest struct {
	Heading    string                `form:"heading" validate:"required,max=200,min=1"`
	Subheading string                `form:"subheading" validate:"required,max=1000,min=1"`
	Thumbnail  *multipart.FileHeader `form:"thumbnail" validate:"required"`
}
