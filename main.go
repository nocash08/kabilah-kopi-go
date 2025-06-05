package main

import (
	"backend/app"
	"backend/config"
	"backend/exception"
	"backend/helper"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize core dependencies
	db := app.NewDB()
	validate := validator.New()
	config.LoadConfig()

	// Setup dependency injection
	injector := config.NewInjector(db, validate)

	// Setup router
	router := config.NewRouter(injector)

	// Start server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	router.PanicHandler = exception.ErrorHandler

	fmt.Println("Server is running on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
