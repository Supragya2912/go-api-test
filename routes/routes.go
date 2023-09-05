package routes

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, userHandler handlers.userHandler) {
	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUsers)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
}
