package repositories

import (
	"github.com/labstack/echo"
	"dome.cloud/secureapi/models"
)

func GetUser(c echo.Context) (interface{}, error) {
	user := models.User{}
	user.CreateUser("Panupong")
	return user, nil
}
