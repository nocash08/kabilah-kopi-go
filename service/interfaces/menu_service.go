package interfaces

import (
	"backend/model/dto/menudto"
	"context"
)

type MenuService interface {
	Create(ctx context.Context, request menudto.MenuCreateRequest) menudto.MenuResponse
	Update(ctx context.Context, request menudto.MenuUpdateRequest) menudto.MenuResponse
	Delete(ctx context.Context, menuId uint)
	FindById(ctx context.Context, menuId uint) menudto.MenuResponse
	FindAll(ctx context.Context) []menudto.MenuResponse
}
