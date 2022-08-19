package routes

import (
	"net/http"

	"github.com/agutama/echo-rest/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/company", controllers.FetchAllCompany)
	e.GET("/companybyid", controllers.FetchCompanyByID)
	e.GET("/companyinid", controllers.FetchCompanyInID)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
