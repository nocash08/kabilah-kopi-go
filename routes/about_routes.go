package routes

import (
	"backend/controller/aboutcontroller"

	"github.com/julienschmidt/httprouter"
)

func SetupAboutRoutes(router *httprouter.Router, aboutController aboutcontroller.AboutController) {
	router.GET("/api/abouts", aboutController.FindAll)
	router.POST("/api/abouts", aboutController.Create)
	router.PUT("/api/abouts/:aboutId", aboutController.Update)
	router.GET("/api/abouts/:aboutId", aboutController.FindById)
}
