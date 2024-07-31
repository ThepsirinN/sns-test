package main

import (
	"net/http"
	"sns-barko/utility/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.NewResponse[any](response.CODE_SUCCESS, "service is running!", nil))
	})

	e.GET("/ready", func(c echo.Context) error {
		return c.JSON(http.StatusOK, response.NewResponse[any](response.CODE_SUCCESS, "service is ready!", nil))
	})

	apiV1 := e.Group("/api/v1")
	v1Group(apiV1)
	return e
}

func v1Group(g *echo.Group) {
	return
}
