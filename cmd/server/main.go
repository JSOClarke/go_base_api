package main

import (
	"base_crud_api/internals/handlers"
	"base_crud_api/internals/pkg/db"
	"base_crud_api/internals/repository"
	"base_crud_api/internals/router"
	"base_crud_api/internals/services"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Should have a struct with all of the handlers so we can pass them all in at once.
func ServerSetup(appHandlers *handlers.AppHandler) *gin.Engine {
	r := gin.Default()
	router.RegisterRoutes(r, appHandlers)
	return r
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(".Env could not be loaded", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
	)

	database := db.Connect(connStr)
	defer database.Close()
	userRepo := repository.NewRepoService(database)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	appHandlers := &handlers.AppHandler{User: userHandler}
	r := ServerSetup(appHandlers)

	r.Run(":4001")
}
