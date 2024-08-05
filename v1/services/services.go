package services

import (
	"context"
	"sns-barko/config"
	"sns-barko/utility/logger"
	"sns-barko/v1/models"
)

type repoV1Interface interface {
	AutoMigrate(ctx context.Context) error
	CreateUser(ctx context.Context, model models.User) error
	ReadUsersByEmail(ctx context.Context, model *models.User) error
	ReadUsersById(ctx context.Context, model *models.User) error
	FindUsersByEmail(ctx context.Context, email string, userId int32, model *[]models.User) error
	UpdateUsers(ctx context.Context, model models.User) error
	DeleteUserById(ctx context.Context, model models.User) error
}

type serviceV1 struct {
	userRepo repoV1Interface
	secret   *config.Secret
}

func New(ctx context.Context, userRepo repoV1Interface, secret *config.Secret) *serviceV1 {
	svc := serviceV1{userRepo, secret}
	err := svc.userRepo.AutoMigrate(ctx)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	return &svc
}
