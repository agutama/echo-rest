package routes

import (
	"net/http"

	"github.com/agutama/echo-rest/controllers"
	"github.com/agutama/echo-rest/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo!")
	})

	e.GET("/pegawai", controllers.FetchAllPegawai, middleware.IsAuthenticated)

	e.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("/login", controllers.CheckLogin)

	return e
}
