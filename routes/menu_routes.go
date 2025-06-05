package routes

import (
	"backend/controller/menucontroller"

	"github.com/julienschmidt/httprouter"
)

func SetupMenuRoutes(router *httprouter.Router, menuController menucontroller.MenuController) {
	router.GET("/api/menus", menuController.FindAll)
	router.POST("/api/menus", menuController.Create)
	router.PUT("/api/menus/:menuId", menuController.Update)
	router.DELETE("/api/menus/:menuId", menuController.Delete)
	router.GET("/api/menus/:menuId", menuController.FindById)
}
