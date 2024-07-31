package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sns-barko/config"
	"sns-barko/utility/logger"
	"time"

	"github.com/labstack/echo/v4"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

func main() {
	bgCtx := context.Background()
	cfg, secret := config.InitConfig()
	logger.InitLogger(cfg)
	ctx, stop := signal.NotifyContext(bgCtx, os.Interrupt, os.Kill)
	defer stop()

	routerV1 := initRouter()
	go initEcho(ctx, routerV1, secret)
	// e.GET("/api-doc/*", echoSwagger.WrapHandler)

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	shutdown(ctx, routerV1)

}

func initEcho(ctx context.Context, e *echo.Echo, secret *config.Secret) {
	if err := e.Start(fmt.Sprint(":", secret.App.Port)); err != nil && err != http.ErrServerClosed {
		logger.Error(ctx, err)
	}
}

func shutdown(ctx context.Context, e *echo.Echo) {
	if err := e.Shutdown(ctx); err != nil {
		logger.Error(ctx, err)
	}
}
