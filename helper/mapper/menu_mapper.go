package mapper

import (
	"backend/model/domain"
	"backend/model/dto/menudto"
)

func ToMenuResponse(menu domain.Menu) menudto.MenuResponse {
	return menudto.MenuResponse{
		Id:         menu.Id,
		Heading:    menu.Heading,
		Subheading: menu.Subheading,
		Thumbnail:  menu.Thumbnail,
	}
}

func ToMenuResponses(menus []domain.Menu) []menudto.MenuResponse {
	var menuResponses []menudto.MenuResponse
	for _, menu := range menus {
		menuResponses = append(menuResponses, ToMenuResponse(menu))
	}
	return menuResponses
}
