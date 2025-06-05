package interfaces

import (
	"backend/model/dto/eventdto"
	"context"
)

type EventService interface {
	Create(ctx context.Context, request eventdto.EventCreateRequest) eventdto.EventResponse
	Update(ctx context.Context, request eventdto.EventUpdateRequest) eventdto.EventResponse
	Delete(ctx context.Context, eventId uint)
	FindById(ctx context.Context, eventId uint) eventdto.EventResponse
	FindAll(ctx context.Context) []eventdto.EventResponse
}
