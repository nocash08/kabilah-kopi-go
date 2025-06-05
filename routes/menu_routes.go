package routes

import (
	"backend/controller/menucontroller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupMenuRoutes(router *httprouter.Router, menuController menucontroller.MenuController, adminMiddleware func(http.Handler) http.Handler) {
	// Public routes (no authentication needed)
	router.GET("/api/menus", menuController.FindAll)
	router.GET("/api/menus/:menuId", menuController.FindById)

	// Protected routes (admin only)
	router.Handler("POST", "/api/menus", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		menuController.Create(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
	router.Handler("PUT", "/api/menus/:menuId", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		menuController.Update(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
	router.Handler("DELETE", "/api/menus/:menuId", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		menuController.Delete(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
}
