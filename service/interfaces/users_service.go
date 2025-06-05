package interfaces

import (
	"backend/model/domain"
	"backend/model/dto/authdto"
	"backend/model/dto/usersdto"
	"context"
	"net/http"
)

type UsersService interface {
	Register(ctx context.Context, request usersdto.UsersRegisterRequest) domain.Users
	Login(ctx context.Context, w http.ResponseWriter, request usersdto.UsersLoginRequest) (domain.Users, *authdto.TokenResponse)
	Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error
}
