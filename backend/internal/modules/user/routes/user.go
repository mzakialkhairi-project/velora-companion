// Package routes provides user route definitions.
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mzakiaklhairi/velora/internal/modules/user/handler"
)

// RegisterRoutes registers the user routes
func RegisterRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {
	// User routes will be registered here when ready
	// Example:
	// users := router.Group("/users")
	// {
	//     users.POST("", userHandler.Create)
	//     users.GET("", userHandler.List)
	//     users.GET("/:id", userHandler.GetByID)
	//     users.PUT("/:id", userHandler.Update)
	//     users.DELETE("/:id", userHandler.Delete)
	// }
}
