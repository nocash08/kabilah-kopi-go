package implementation

import (
	"backend/helper"
	"backend/model/domain"
	"backend/model/dto/authdto"
	"backend/model/dto/usersdto"
	"backend/repository/usersrepository"
	"backend/service/interfaces"
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UsersServiceImpl struct {
	UsersRepository usersrepository.UsersRepository
	DB              *sql.DB
	Validator       *validator.Validate
	JWTSecret       string
}

func NewUsersService(
	usersRepository usersrepository.UsersRepository,
	_ interface{}, // Keep this parameter to maintain compatibility
	DB *sql.DB,
	validator *validator.Validate,
	jwtSecret string,
) interfaces.UsersService {
	return &UsersServiceImpl{
		UsersRepository: usersRepository,
		DB:              DB,
		Validator:       validator,
		JWTSecret:       jwtSecret,
	}
}

func (service *UsersServiceImpl) Login(ctx context.Context, w http.ResponseWriter, request usersdto.UsersLoginRequest) (domain.Users, *authdto.TokenResponse) {
	err := service.Validator.Struct(request)
	if err != nil {
		panic(err)
	}

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
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(request.Password))
	if err != nil {
		panic(errors.New("invalid username or password"))
	}

	// Generate JWT token
	tokens, err := helper.GenerateTokens(user.Id, user.Username, user.IsAdmin, service.JWTSecret)
	if err != nil {
		panic(err)
	}

	// Set access token cookie
	helper.SetTokenCookie(w, helper.AccessTokenCookie, tokens.AccessToken, helper.AccessTokenDuration)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	return user, &authdto.TokenResponse{
		AccessToken: tokens.AccessToken,
		TokenType:   "Bearer",
		ExpiresIn:   int(helper.AccessTokenDuration.Seconds()),
	}
}

func (service *UsersServiceImpl) Logout(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return errors.New("authorization header is required")
	}

	// Check if the header has the Bearer prefix
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return errors.New("invalid token format. Use 'Bearer <token>'")
	}

	// Extract the token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == "" {
		return errors.New("token is required")
	}

	// Parse the token to get expiration time
	token, err := jwt.ParseWithClaims(tokenString, &helper.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(service.JWTSecret), nil
	})
	if err != nil {
		return errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(*helper.JWTClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	// Add token to invalidated tokens list
	helper.AddToInvalidatedTokens(tokenString, claims.ExpiresAt.Time)

	// Clear the access token cookie
	helper.ClearTokenCookie(w, helper.AccessTokenCookie)
	return nil
}

func (service *UsersServiceImpl) Register(ctx context.Context, request usersdto.UsersRegisterRequest) domain.Users {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Check if username already exists
	_, _, err = service.UsersRepository.FindByUsername(ctx, tx, request.Username)
	if err == nil {
		panic(errors.New("username already exists"))
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	// Check if this is the first user
	userCount := service.UsersRepository.CountUsers(ctx, tx)
	isAdmin := userCount == 0 // Make first user an admin

	user := domain.Users{
		Username: request.Username,
		Password: string(hashedPassword),
		IsAdmin:  isAdmin,
	}

	return service.UsersRepository.Create(ctx, tx, user)
}
