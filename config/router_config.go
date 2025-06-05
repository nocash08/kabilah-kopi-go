package config

import (
	"backend/routes"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(injector *Injector) *httprouter.Router {
	router := httprouter.New()

	// Setup all routes
	routes.SetupMenuRoutes(router, *injector.MenuController)
	routes.SetupAboutRoutes(router, *injector.AboutController)
	routes.SetupEventRoutes(router, *injector.EventController)
	routes.SetupUsersRoutes(router, *injector.UsersController)

	return router
}
