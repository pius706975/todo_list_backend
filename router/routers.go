package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterApp() *echo.Echo {
	
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is To Do List App")
	})

	return e
}