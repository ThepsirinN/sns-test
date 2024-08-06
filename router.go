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
	// e.GET("/swagger/*", echoSwagger.WrapHandler)
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

	friendV1 := apiV1.Group("/friend")
	v1FriendAuthGroup(friendV1, handlersV1, jwtSecret)

	postV1 := apiV1.Group("/post")
	v1PostAuthGroup(postV1, handlersV1, jwtSecret)

	likeV1 := apiV1.Group("/like")
	v1LikeAuthGroup(likeV1, handlersV1, jwtSecret)

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

func v1FriendAuthGroup(g *echo.Group, handlersV1 *handlers.HandlersV1, jwtSecret string) {
	g.Use(auth.MiddleWareAuth(jwtSecret))
	g.POST("/create", handlersV1.CreateFriendRequest)
	g.GET("/all-friend-request", handlersV1.GetAllFriendRequest) // can separate for sending and receive
	g.GET("/list-all-friend", handlersV1.ListFriend)
	g.PATCH("/update-status", handlersV1.UpdateFriendRequestStatus)
	g.DELETE("/delete", handlersV1.DeleteFriend)
}

func v1PostAuthGroup(g *echo.Group, handlersV1 *handlers.HandlersV1, jwtSecret string) {
	g.Use(auth.MiddleWareAuth(jwtSecret))
	g.POST("/create", handlersV1.CreatePost)
	g.GET("/all-posts", handlersV1.ListAllPostFromUser) // can separate for sending and receive
	g.GET("/post/:id", handlersV1.ReadPostByPostId)
	g.PATCH("/update", handlersV1.UpdatePostData)
	g.DELETE("/delete", handlersV1.DeletePost)
}

func v1LikeAuthGroup(g *echo.Group, handlersV1 *handlers.HandlersV1, jwtSecret string) {
	g.Use(auth.MiddleWareAuth(jwtSecret))
	g.POST("/add", handlersV1.AddLike)
	g.DELETE("/delete", handlersV1.DeleteLike)
}
