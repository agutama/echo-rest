package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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

func FetchAllPegawaiByID(c echo.Context) error {

	id := c.FormValue("id")

	conv_id, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := models.FetchPegawaiByID(conv_id)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
