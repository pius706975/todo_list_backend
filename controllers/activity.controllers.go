package controllers

import (
	"net/http"
	"pius/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

func AddActivityCTRL(c echo.Context) error {

	title := c.FormValue("title")
	email := c.FormValue("email")

	if govalidator.IsNull(title) {
		res := models.Response{
			Status: http.StatusBadRequest,
			Message: "Title cannot be empty",
			Data: "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	result, err := models.AddActivity(title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteActivityCTRL(c echo.Context) error {

	ID := c.Param("id")

	idINT, err := strconv.Atoi(ID)
	if err != nil {
		res := models.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	getData, err := models.DeleteActivity(idINT)
	if err != nil {
		return c.JSON(getData.Status, getData)
	}

	return c.JSON(getData.Status, getData)
}

func UpdateActivityCTRL(c echo.Context) error {

	ID := c.Param("id")
	title := c.FormValue("title")
	email := c.FormValue("email")

	idINT, err := strconv.Atoi(ID)
	if err != nil {
		res := models.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	updatedData, err := models.UpdateActivity(idINT, title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(updatedData.Status, updatedData)
}

func GetAllActivitiesCTRL(c echo.Context) error {

	result, err := models.GetAllActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetByIDCTRL(c echo.Context) error {

	ID := c.Param("id")

	idINT, err := strconv.Atoi(ID)
	if err != nil {
		res := models.Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid ID",
			Data:    "",
		}
		return c.JSON(http.StatusBadRequest, res)
	}

	activity, err := models.GetByID(idINT)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": activity,
	})

}
