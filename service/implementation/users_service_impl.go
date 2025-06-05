package implementation

import (
	"backend/helper"
	"backend/model/dto/authdto"
	"backend/model/dto/usersdto"
	"backend/repository/usersrepository"
	"backend/service/interfaces"
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository usersrepository.UsersRepository
	DB              *sql.DB
	Validator       *validator.Validate
	JWTSecret       string
}

func NewUsersService(usersRepository usersrepository.UsersRepository, DB *sql.DB, validator *validator.Validate, jwtSecret string) interfaces.UsersService {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
		DB:              DB,
		Validator:       validator,
		JWTSecret:       jwtSecret,
	}
}

func (service *UsersServiceImpl) Login(ctx context.Context, w http.ResponseWriter, request usersdto.UsersLoginRequest) (usersdto.UsersResponse, *authdto.TokenResponse) {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	user, hashedPassword, err := service.UsersRepository.FindByUsername(ctx, tx, request.Username)
	if err != nil {
		panic(errors.New("invalid username or password"))
	}

	// Verify password
	err = helper.VerifyPassword(hashedPassword, request.Password)
	if err != nil {
		panic(errors.New("invalid username or password"))
	}

	// Generate JWT tokens
	tokens, err := helper.GenerateTokens(user.Id, user.Username, user.IsAdmin, service.JWTSecret)
	if err != nil {
		panic(err)
	}

	// Set cookies
	helper.SetTokenCookie(w, helper.AccessTokenCookie, tokens.AccessToken, helper.AccessTokenDuration)
	helper.SetTokenCookie(w, helper.RefreshTokenCookie, tokens.RefreshToken, helper.RefreshTokenDuration)

	tx.Commit()

	return user, &authdto.TokenResponse{
		AccessToken: tokens.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(helper.AccessTokenDuration.Seconds()),
	}
}

func (service *UsersServiceImpl) RefreshToken(ctx context.Context, w http.ResponseWriter, r *http.Request) (*authdto.TokenResponse, error) {
	// Get refresh token from cookie
	cookie, err := r.Cookie(helper.RefreshTokenCookie)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	// Validate refresh token
	claims, err := helper.ValidateToken(cookie.Value, service.JWTSecret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Generate new tokens
	tokens, err := helper.GenerateTokens(claims.UserID, claims.Username, claims.IsAdmin, service.JWTSecret)
	if err != nil {
		return nil, err
	}

	// Set new cookies
	helper.SetTokenCookie(w, helper.AccessTokenCookie, tokens.AccessToken, helper.AccessTokenDuration)
	helper.SetTokenCookie(w, helper.RefreshTokenCookie, tokens.RefreshToken, helper.RefreshTokenDuration)

	return &authdto.TokenResponse{
		AccessToken: tokens.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(helper.AccessTokenDuration.Seconds()),
	}, nil
}

func (service *UsersServiceImpl) Logout(ctx context.Context, w http.ResponseWriter) error {
	// Clear both tokens
	helper.ClearTokenCookie(w, helper.AccessTokenCookie)
	helper.ClearTokenCookie(w, helper.RefreshTokenCookie)
	return nil
}
