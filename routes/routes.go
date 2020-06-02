package routes

import (
	"github.com/labstack/echo"
	"dome.cloud/secureapi/actions/handlers"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/user", handlers.GetUser)
	return e
}
