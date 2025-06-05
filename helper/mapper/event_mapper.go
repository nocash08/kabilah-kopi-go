package mapper

import (
	"backend/model/domain"
	"backend/model/dto/eventdto"
)

func ToEventResponse(event domain.Event) eventdto.EventResponse {
	return eventdto.EventResponse{
		Id:         event.Id,
		Heading:    event.Heading,
		Subheading: event.Subheading,
		Thumbnail:  event.Thumbnail,
	}
}

func ToEventResponses(events []domain.Event) []eventdto.EventResponse {
	var eventResponses []eventdto.EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}
	return eventResponses
}
