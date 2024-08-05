package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sns-barko/config"
	"sns-barko/utility/logger"
	"sns-barko/utility/tracer"
	"sns-barko/v1/handlers"
	"sns-barko/v1/repositories"
	"sns-barko/v1/services"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger/example/docs"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	bgCtx := context.Background()
	cfg, secret := config.InitConfig()
	logger.InitLogger(cfg)
	defer logger.Sync()

	db := initDatabse(bgCtx, cfg, secret)
	repoV1 := repositories.New(db)
	svcV1 := services.New(bgCtx, repoV1, secret)
	handlersV1 := handlers.New(svcV1)

	ctx, stop := signal.NotifyContext(bgCtx, os.Interrupt, os.Kill)
	defer stop()
	tp := tracer.InitTraceProvider(ctx, cfg.Log.Env, "sns-test")
	defer func() {
		if err := tp.Shutdown(ctx); err != nil {
			logger.Fatal(ctx, err)
		}
	}()

	routerV1 := initRouter(cfg, secret.User.JWT.Secret, handlersV1)
	go initEcho(ctx, routerV1, secret)
	// e.GET("/api-doc/*", echoSwagger.WrapHandler)

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logger.Info(ctx, "service is shutting down")
	shutdown(ctx, routerV1)
	log.Println("service is shutted down")

}

func initEcho(ctx context.Context, e *echo.Echo, secret *config.Secret) {
	if err := e.Start(fmt.Sprint(":", secret.App.Port)); err != nil && err != http.ErrServerClosed {
		logger.Fatal(ctx, err)
	}
}

func shutdown(ctx context.Context, e *echo.Echo) {
	if err := e.Shutdown(ctx); err != nil {
		logger.Error(ctx, err)
	}
}

func initDatabse(ctx context.Context, cfg *config.Config, secret *config.Secret) *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		secret.Database.User,
		secret.Database.Password,
		secret.Database.Host,
		secret.Database.Port,
		cfg.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NowFunc: func() time.Time { return time.Now().Local() }})
	if err != nil {
		logger.Fatal(ctx, err)
	}

	// config connection pools
	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifeTime)

	logger.Info(ctx, "Database is running!",
		zap.Any("db_stats", sqlDB.Stats()),
		zap.Int("max_idle_conns", cfg.Database.MaxIdleConns),
		zap.Int("max_open_connection", cfg.Database.MaxOpenConns),
		zap.Duration("max_life_time_minutes", cfg.Database.ConnMaxLifeTime),
		zap.String("max_life_time_minutes_string", cfg.Database.ConnMaxLifeTime.String()),
	)
	return db
}
