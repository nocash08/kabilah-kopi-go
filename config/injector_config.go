package config

import (
	"database/sql"

	"backend/controller/aboutcontroller"
	"backend/controller/eventcontroller"
	"backend/controller/menucontroller"
	"backend/controller/userscontroller"
	"backend/repository/aboutrepository"
	"backend/repository/eventrepository"
	"backend/repository/menurepository"
	"backend/repository/usersrepository"
	"backend/service/implementation"

	"github.com/go-playground/validator/v10"
)

type Injector struct {
	MenuController  *menucontroller.MenuController
	AboutController *aboutcontroller.AboutController
	EventController *eventcontroller.EventController
	UsersController *userscontroller.UsersController
}

func NewInjector(db *sql.DB, validate *validator.Validate) *Injector {
	// Menu dependencies
	menuRepository := menurepository.NewMenuRepository()
	menuService := implementation.NewMenuService(menuRepository, db, validate)
	menuController := menucontroller.NewMenuController(menuService)

	// About dependencies
	aboutRepository := aboutrepository.NewAboutRepository()
	aboutService := implementation.NewAboutService(aboutRepository, db, validate)
	aboutController := aboutcontroller.NewAboutController(aboutService)

	// Event dependencies
	eventRepository := eventrepository.NewEventRepository()
	eventService := implementation.NewEventService(eventRepository, db, validate)
	eventController := eventcontroller.NewEventController(eventService)

	// Users dependencies
	usersRepository := usersrepository.NewUsersRepository()
	usersService := implementation.NewUsersService(usersRepository, nil, db, validate, AppConfig.JWTSecret)
	usersController := userscontroller.NewUsersController(usersService)

	return &Injector{
		MenuController:  &menuController,
		AboutController: &aboutController,
		EventController: &eventController,
		UsersController: &usersController,
	}
}
