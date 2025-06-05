package main

import (
	"backend/app"
	"backend/config"
	"backend/helper"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize database
	db := app.NewDB()
	defer db.Close()

	// Initialize validator
	validate := validator.New()

	// Initialize token cleanup
	helper.InitTokenCleanup()

	// Initialize dependency injection
	injector := config.NewInjector(db, validate)

	// Initialize router
	router := config.NewRouter(injector)

	// Setup CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Create server with CORS handler
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: corsHandler.Handler(router),
	}

	// Start server
	fmt.Println("Server is running on http://localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
