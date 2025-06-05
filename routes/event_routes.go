package routes

import (
	"backend/controller/eventcontroller"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupEventRoutes(router *httprouter.Router, eventController eventcontroller.EventController, adminMiddleware func(http.Handler) http.Handler) {
	// Public routes (no authentication needed)
	router.GET("/api/events", eventController.FindAll)
	router.GET("/api/events/:eventId", eventController.FindById)

	// Protected routes (admin only)
	router.Handler("POST", "/api/events", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eventController.Create(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
	router.Handler("PUT", "/api/events/:eventId", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eventController.Update(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
	router.Handler("DELETE", "/api/events/:eventId", adminMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eventController.Delete(w, r, httprouter.ParamsFromContext(r.Context()))
	})))
}
