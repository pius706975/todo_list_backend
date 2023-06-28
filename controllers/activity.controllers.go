package controllers

import (
	"net/http"
	"pius/models"

	"github.com/labstack/echo/v4"
)

func GetAllActivitiesCTRL(c echo.Context) error {
	
	result, err := models.GetAllActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetByIDCTRL(c echo.Context) error {

	// Get ID from request parameters
	ID := c.Param("id")

	activity, err := models.GetByID(ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": activity,
	})


}

func AddActivityCTRL(c echo.Context) error {

	title := c.FormValue("title")
	email := c.FormValue("email")

	result, err := models.AddActivity(title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}