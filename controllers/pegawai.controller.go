package controllers

import (
	"net/http"

	"github.com/agutama/echo-rest/models"
	"github.com/labstack/echo"
)

func FetchAllPegawai(c echo.Context) error {
	result, err := models.FectAllPegawai()

	if err != nil {
		return c.JSON(http.StatusOK, "Hello, this is echo")
	}

	return c.JSON(http.StatusOK, result)
}
