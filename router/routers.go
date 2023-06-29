package router

import (
	"net/http"
	"pius/controllers"

	"github.com/labstack/echo/v4"
)

func RouterApp() *echo.Echo {
	
	e := echo.New()

	group := e.Group("")

	activity := group.Group("activity-groups")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is To Do List App")
	})

	// Activity Controllers
	activity.GET("/", controllers.GetAllActivitiesCTRL)
	activity.GET("/:id", controllers.GetByIDCTRL)
	activity.POST("/create", controllers.AddActivityCTRL)
	activity.DELETE("/delete/:id", controllers.DeleteActivityCTRL)
	activity.PATCH("/update/:id", controllers.UpdateActivityCTRL)

	return e
}