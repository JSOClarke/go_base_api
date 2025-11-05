package router

import (
	"base_crud_api/internals/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, appHandlers *handlers.AppHandler) {
	v1 := r.Group("/api/v1")
	{
		// v1.GET("/health", app)
		users := v1.Group("/users")
		{
			users.POST("/signup", appHandlers.User.SignUpUser)
			users.POST("/login", appHandlers.User.LoginUser)
		}

	}
}
