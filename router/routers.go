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
	todo := group.Group("todo-items")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is To Do List App")
	})

	// Activity Controllers
	activity.GET("", controllers.GetAllActivitiesCTRL)
	activity.GET("/:id", controllers.GetByIDCTRL)
	activity.POST("", controllers.AddActivityCTRL)
	activity.DELETE("/:id", controllers.DeleteActivityCTRL)
	activity.PATCH("/:id", controllers.UpdateActivityCTRL)

	// To-Do Controllers
	todo.GET("/", controllers.GetAllTodoItemsCTRL)
	todo.GET("", controllers.GetTodoItemsByGroupCTRL)
	todo.GET("/:id", controllers.GetTodoByIDCTRL)
	todo.POST("/create", controllers.AddTodoItemCTRL)
	todo.DELETE("/:id", controllers.DeleteTodoItemCTRL)
	todo.PATCH("/:id", controllers.UpdateTodoItemCTRL)

	return e
}