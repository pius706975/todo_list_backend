package controllers

import (
	"net/http"
	"pius/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

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