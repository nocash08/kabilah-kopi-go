package routes

import (
	"backend/controller/userscontroller"
	"backend/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupUsersRoutes(router *httprouter.Router, usersController userscontroller.UsersController) {
	router.POST("/api/register", usersController.Register)
	router.POST("/api/login", usersController.Login)

	// Protected route - requires valid JWT token
	router.Handler("POST", "/api/logout", middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		usersController.Logout(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
}
