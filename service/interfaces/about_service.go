package interfaces

import (
	"backend/model/dto/aboutdto"
	"context"
)

type AboutService interface {
	Create(ctx context.Context, request aboutdto.AboutCreateRequest) aboutdto.AboutResponse
	Update(ctx context.Context, request aboutdto.AboutUpdateRequest) aboutdto.AboutResponse
	FindById(ctx context.Context, aboutId uint) aboutdto.AboutResponse
	FindAll(ctx context.Context) []aboutdto.AboutResponse
}
