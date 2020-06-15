package repositories

import (
	"dome.cloud/secureapi/models"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	user.CreateUser("Panupong")
	return user, nil
}
