package routes

import (
	"backend/controller/userscontroller"

	"github.com/julienschmidt/httprouter"
)

func SetupUsersRoutes(router *httprouter.Router, usersController userscontroller.UsersController) {
	router.POST("/api/login", usersController.Login)
	router.POST("/api/refresh-token", usersController.RefreshToken)
	router.POST("/api/logout", usersController.Logout)
}
