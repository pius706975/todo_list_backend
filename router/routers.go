package router

import (
	"net/http"
	"pius/controllers"

	"github.com/labstack/echo/v4"
)

func RouterApp() *echo.Echo {
	
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is To Do List App")
	})

	// Activity Controllers
	e.GET("/activities", controllers.GetAllActivitiesCTRL)
	e.GET("/activities/:id", controllers.GetByIDCTRL)
	e.POST("/add_activity", controllers.AddActivityCTRL)

	return e
}