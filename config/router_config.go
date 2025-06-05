package config

import (
	"backend/middleware"
	"backend/routes"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(injector *Injector) *httprouter.Router {
	router := httprouter.New()

	// Create middleware chain
	jwtAndAdminMiddleware := func(next http.Handler) http.Handler {
		return middleware.JWTMiddleware(middleware.AdminMiddleware(next))
	}

	// Setup all routes
	routes.SetupMenuRoutes(router, *injector.MenuController, jwtAndAdminMiddleware)
	routes.SetupAboutRoutes(router, *injector.AboutController, jwtAndAdminMiddleware)
	routes.SetupEventRoutes(router, *injector.EventController, jwtAndAdminMiddleware)
	routes.SetupUsersRoutes(router, *injector.UsersController)

	return router
}
