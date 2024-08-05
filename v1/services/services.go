package services

import (
	"context"
	"sns-barko/config"
	"sns-barko/utility/logger"
	"sns-barko/v1/entities"
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

	CreateFriendRequest(ctx context.Context, model models.Friend) error
	GetAllFriendRequest(ctx context.Context, id int32, model *[]entities.GetAllFriendRequestResponse) error
	ListFriend(ctx context.Context, id int32, model *[]entities.ListFriendQuery) error
	UpdateFriendRequestStatus(ctx context.Context, model models.Friend) error
	DeleteFriend(ctx context.Context, model models.Friend) error
}

type cacheV1 interface {
	ClearMultipleCacheKey(ctx context.Context, key ...string) error
	GetDataFromCache(ctx context.Context, key string) ([]byte, error)
	SetDataToCache(ctx context.Context, key string, value []byte) error
}

type serviceV1 struct {
	repoV1  repoV1Interface
	cacheV1 cacheV1
	secret  *config.Secret
}

func New(ctx context.Context, repoV1 repoV1Interface, cacheV1 cacheV1, secret *config.Secret) *serviceV1 {
	svc := serviceV1{repoV1, cacheV1, secret}
	err := svc.repoV1.AutoMigrate(ctx)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	return &svc
}
