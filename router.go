package main

import (
	"net/http"
	"sns-barko/config"
	"sns-barko/utility/logger"
	"sns-barko/utility/response"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/propagation"
)

func initRouter(cfg *config.Config) *echo.Echo {
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
	propagator := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)

	apiV1.Use(otelecho.Middleware(cfg.App.Name, otelecho.WithPropagators(propagator)))
	v1Group(apiV1)
	return e
}

func v1Group(g *echo.Group) {
	g.GET("/test", func(c echo.Context) error { logger.Info(c.Request().Context(), "test"); return nil })
}
