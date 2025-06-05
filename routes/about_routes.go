package routes

import (
	"backend/controller/aboutcontroller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupAboutRoutes(router *httprouter.Router, aboutController aboutcontroller.AboutController, adminMiddleware func(http.Handler) http.Handler) {
	// Public routes (no authentication needed)
	router.GET("/api/abouts", aboutController.FindAll)
	router.GET("/api/abouts/:aboutId", aboutController.FindById)

	// Protected routes (admin only)
	router.Handler("POST", "/api/abouts", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aboutController.Create(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
	router.Handler("PUT", "/api/abouts/:aboutId", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		aboutController.Update(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
}
