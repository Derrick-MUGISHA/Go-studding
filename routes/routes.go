package routes

import (
	"go-api/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up the routes for the API.
// It takes in a gin.Engine and adds a GET route for /users
// that calls the handlers.GetUsers function.
func SetupRoutes(router *gin.Engine) {
	router.GET("/users", handlers.GetUsers)
}
