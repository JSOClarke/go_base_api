package router

import (
	"base_crud_api/internals/handlers"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusHandler() gin.HandlerFunc {
    h := promhttp.Handler()

    return func(c *gin.Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}


func RegisterRoutes(r *gin.Engine, appHandlers *handlers.AppHandler) {
	
			r.GET("/metrics",PrometheusHandler())

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
