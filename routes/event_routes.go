package routes

import (
	"backend/controller/eventcontroller"

	"github.com/julienschmidt/httprouter"
)

func SetupEventRoutes(router *httprouter.Router, eventController eventcontroller.EventController) {
	router.GET("/api/events", eventController.FindAll)
	router.POST("/api/events", eventController.Create)
	router.PUT("/api/events/:eventId", eventController.Update)
	router.DELETE("/api/events/:eventId", eventController.Delete)
	router.GET("/api/events/:eventId", eventController.FindById)
}
