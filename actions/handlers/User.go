package handlers

import (
	"net/http"

	"dome.cloud/secureapi/actions/repositories"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	user, err := repositories.GetUser(c)
	if err != nil {
		return c.JSON(404, "")
	}
	return c.JSON(http.StatusOK, user)
}
