package routes

import (
	"dome.cloud/secureapi/actions/handlers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/user", handlers.GetUser)
	return e
}
