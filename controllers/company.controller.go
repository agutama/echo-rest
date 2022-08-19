package controllers

import (
	"net/http"

	"github.com/agutama/echo-rest/models"
	"github.com/labstack/echo"
)

func FetchAllCompany(c echo.Context) error {
	result, err := models.FectAllCompany()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func FetchCompanyByID(c echo.Context) error {

	id := c.FormValue("id")

	result, err := models.GetCompanyByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func FetchCompanyInID(c echo.Context) error {

	id := c.FormValue("id")

	result, err := models.GetCompanyInID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
