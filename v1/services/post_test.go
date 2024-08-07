package services

import (
	"context"
	"errors"
	"sns-barko/v1/entities"

	"github.com/stretchr/testify/mock"
)

func (t *ServTestSuite) Test_CreatePost_DB_Error() {
	t.wantErr = true
	req := entities.CreatePostRequest{}

	t.mockRepoV1.EXPECT().CreatePost(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.CreatePost(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.CreatePost() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_CreatePost_DB_Success() {
	t.wantErr = false
	req := entities.CreatePostRequest{}

	t.mockRepoV1.EXPECT().CreatePost(t.ctx, mock.Anything).Return(nil).Once()

	err := t.svc.CreatePost(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.CreatePost() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_ListAllPostFromUser_DB_Error() {
	t.wantErr = true
	req := entities.ListAllPostFromUserRequest{}
	entity := []entities.ListAllPostFromUserResponse{}

	t.mockRepoV1.EXPECT().ListAllPostFromUser(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.ListAllPostFromUser(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListAllPostFromUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_ListAllPostFromUser_DB_Success() {
	t.wantErr = false
	req := entities.ListAllPostFromUserRequest{}
	entity := []entities.ListAllPostFromUserResponse{}

	mockData := []entities.ListAllPostFromUserResponse{}
	t.mockRepoV1.EXPECT().ListAllPostFromUser(t.ctx, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, i int32, lapfur *[]entities.ListAllPostFromUserResponse) error {
		*lapfur = mockData
		return nil
	}).Once()

	err := t.svc.ListAllPostFromUser(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListAllPostFromUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

	resp := []entities.ListAllPostFromUserResponse{}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.ListAllPostFromUser() = %v, want %v", entity, resp)
	}
}

func (t *ServTestSuite) Test_ReadPostByPostId_DB_Error() {
	t.wantErr = false
	req := entities.ReadPostByPostIdRequest{}
	entity := entities.ReadPostByPostIdResponse{}

	t.mockRepoV1.EXPECT().ReadPostByPostId(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.ReadPostByPostId(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ReadPostByPostId() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_ReadPostByPostId_DB_Success() {
	t.wantErr = false
	req := entities.ReadPostByPostIdRequest{}
	entity := entities.ReadPostByPostIdResponse{}

	mockData := entities.ReadPostByPostIdResponse{}
	t.mockRepoV1.EXPECT().ReadPostByPostId(t.ctx, mock.Anything).RunAndReturn(func(ctx context.Context, rpbpir *entities.ReadPostByPostIdResponse) error {
		*rpbpir = mockData
		return nil
	}).Once()

	err := t.svc.ReadPostByPostId(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ReadPostByPostId() error = %v, wantErr %v", err, t.wantErr)
		return
	}

	resp := entities.ReadPostByPostIdResponse{}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.ReadPostByPostId() = %v, want %v", entity, resp)
	}
}

func (t *ServTestSuite) Test_UpdatePost_DB_Error() {
	t.wantErr = true
	req := entities.UpdatePostRequest{}

	t.mockRepoV1.EXPECT().UpdatePostData(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.UpdatePostData(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.UpdatePostData() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_UpdatePost_DB_Success() {
	t.wantErr = false
	req := entities.UpdatePostRequest{}

	t.mockRepoV1.EXPECT().UpdatePostData(t.ctx, mock.Anything).Return(nil).Once()

	err := t.svc.UpdatePostData(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.UpdatePostData() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeletePost_DB_Error() {
	t.wantErr = true
	req := entities.DeletePostRequest{}

	t.mockRepoV1.EXPECT().DeletePost(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.DeletePost(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeletePost() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeletePost_DB_Success() {
	t.wantErr = false
	req := entities.DeletePostRequest{}

	t.mockRepoV1.EXPECT().DeletePost(t.ctx, mock.Anything).Return(nil).Once()

	err := t.svc.DeletePost(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeletePost() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}
