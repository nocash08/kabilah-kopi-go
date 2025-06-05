package interfaces

import (
	"backend/model/dto/authdto"
	"backend/model/dto/usersdto"
	"context"
	"net/http"
)

type UsersService interface {
	Login(ctx context.Context, w http.ResponseWriter, request usersdto.UsersLoginRequest) (usersdto.UsersResponse, *authdto.TokenResponse)
	RefreshToken(ctx context.Context, w http.ResponseWriter, r *http.Request) (*authdto.TokenResponse, error)
	Logout(ctx context.Context, w http.ResponseWriter) error
}
