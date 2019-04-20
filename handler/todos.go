package handler

import (
	"github.com/labstack/echo"
	"net/http"

	"github.com/h-hiroki/echo_first/models"
)

func GetTodos(c echo.Context) error {
	return c.JSON(http.StatusOK, models.GetTodos())
}
