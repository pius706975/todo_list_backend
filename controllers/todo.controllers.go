package controllers

import (
	"net/http"
	"pius/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddTodoItemCTRL(c echo.Context) error {
	
	activitiGroupID := c.FormValue("activity_group_id")

	idINT, err := strconv.Atoi(activitiGroupID)
	if err != nil {
		res := models.Response{
			Status: http.StatusBadRequest,
			Message: "Invalid ID",
			Data: "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	title := c.FormValue("title")
	priority := c.FormValue("priority")

	result, err := models.AddTodoItem(idINT, title, priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteTodoItemCTRL(c echo.Context) error {
	
	ID := c.Param("id")

	idINT, err := strconv.Atoi(ID)
	if err != nil {
		res := models.Response{
			Status: http.StatusBadRequest,
			Message: "Invalid ID",
			Data: "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	getData, err := models.DeleteTodoItem(idINT)
	if err != nil {
		return c.JSON(getData.Status, getData)
	}

	return c.JSON(getData.Status, getData)
}

func GetTodoByIDCTRL(c echo.Context) error {
	
	ID := c.Param("id")
	
	idINT, err := strconv.Atoi(ID)
	if err != nil {
		res := models.Response{
			Status: http.StatusBadRequest,
			Message: "Invalid ID",
			Data: "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	todo, err := models.GetTodoByID(idINT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": todo,
	})
}