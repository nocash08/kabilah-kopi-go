package mapper

import (
	"backend/model/domain"
	"backend/model/dto/aboutdto"
)

func ToAboutResponse(about domain.About) aboutdto.AboutResponse {
	return aboutdto.AboutResponse{
		Id:         about.Id,
		Heading:    about.Heading,
		Subheading: about.Subheading,
	}
}

func ToAboutResponses(abouts []domain.About) []aboutdto.AboutResponse {
	var aboutResponses []aboutdto.AboutResponse
	for _, about := range abouts {
		aboutResponses = append(aboutResponses, ToAboutResponse(about))
	}
	return aboutResponses
}
