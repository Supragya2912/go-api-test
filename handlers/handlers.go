package handlers

import (
	"api-golang/config"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
}

func (h *userHandler) CreateUser(c echo.Context) error {

	var user users.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	err := users.CreateUser(config.GetMongoDBClient(), user)
	if err != nil {
		return err
	}

	return c.JSON(201, user)
}
