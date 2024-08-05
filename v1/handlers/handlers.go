package handlers

import (
	"context"
	"sns-barko/v1/entities"
)

type servicesV1Interface interface {
	CreateUser(ctx context.Context, req entities.CreateUserRequest) error
	AuthUser(ctx context.Context, req entities.AuthUserRequest, resp *entities.AuthUserResponse) error
	FindUsersByEmail(ctx context.Context, userId int32, req entities.FindUserByEmailRequest, resp *[]entities.FindUserByEmailResponse) error
	UpdateUser(ctx context.Context, req entities.UpdateUserRequest, resp *entities.UpdateUserResponse) error
	DeleteUser(ctx context.Context, req entities.DeleteUserRequest) error

	CreateFriendRequest(ctx context.Context, req entities.CreateFriendRequestRequest) error
	GetAllFriendRequest(ctx context.Context, req entities.GetAllFriendRequestRequest, resp *[]entities.GetAllFriendRequestResponse) error
	ListFriend(ctx context.Context, req entities.ListFriendRequest, resp *[]entities.ListFriendResponse) error
	UpdateFriendRequestStatus(ctx context.Context, req entities.UpdateFriendRequestStatusRequest) error
	DeleteFriend(ctx context.Context, req entities.DeleteFriendRequestRequest) error
}

type HandlersV1 struct {
	service servicesV1Interface
}

func New(service servicesV1Interface) *HandlersV1 {
	handler := HandlersV1{service}
	return &handler
}
