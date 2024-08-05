package main

import (
	"net/http"
	"sns-barko/config"
	"sns-barko/middleware/auth"
	"sns-barko/utility/response"
	"sns-barko/v1/handlers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/propagation"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}

func initRouter(cfg *config.Config, jwtSecret string, handlersV1 *handlers.HandlersV1) *echo.Echo {
	e := echo.New()
	e.Validator = &customValidator{validator: validator.New()}
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

	userV1 := apiV1.Group("/user")
	v1UserGroup(userV1, handlersV1)
	v1UserAuthGroup(userV1, handlersV1, jwtSecret)

	return e
}

func v1UserGroup(g *echo.Group, handlersV1 *handlers.HandlersV1) {
	g.POST("/register", handlersV1.CreateUser)
	g.POST("/auth", handlersV1.AuthUser)
}

func v1UserAuthGroup(g *echo.Group, handlersV1 *handlers.HandlersV1, jwtSecret string) {
	g.Use(auth.MiddleWareAuth(jwtSecret))
	g.DELETE("/delete", handlersV1.DeleteUser)
	g.GET("/find_user/:email", handlersV1.FindUserByEmail)
	g.PATCH("/update", handlersV1.UpdateUser)
}
