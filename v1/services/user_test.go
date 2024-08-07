package services

import (
	"context"
	"errors"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"

	"github.com/stretchr/testify/mock"
)

func (t *ServTestSuite) Test_CreateUser_GeneratePasswordError() {
	t.wantErr = true
	req := entities.CreateUserRequest{
		Auth: "1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p1234567890p",
	}

	err := t.svc.CreateUser(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.CreateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_CreateUser_GeneratePassword_DB_Error() {
	t.wantErr = true
	req := entities.CreateUserRequest{
		Auth: "asd",
	}
	t.mockRepoV1.EXPECT().CreateUser(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.CreateUser(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.CreateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_CreateUser_GeneratePassword_DB_Success() {
	t.wantErr = false
	req := entities.CreateUserRequest{
		Auth: "asd",
	}
	t.mockRepoV1.EXPECT().CreateUser(t.ctx, mock.Anything).Return(nil).Once()

	err := t.svc.CreateUser(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.CreateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_AuthUser_ReadUser_DB_Error() {
	t.wantErr = true
	req := entities.AuthUserRequest{}
	t.mockRepoV1.EXPECT().ReadUsersByEmail(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.AuthUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.AuthUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_AuthUser_ReadUser_DB_Password_Compare_Error() {
	t.wantErr = true
	req := entities.AuthUserRequest{Password: "123"}
	t.mockRepoV1.EXPECT().ReadUsersByEmail(t.ctx, mock.Anything).Return(nil).Once()

	err := t.svc.AuthUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.AuthUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_AuthUser_ReadUser_DB_Password_Compare_Sign_Token() {
	t.wantErr = false
	req := entities.AuthUserRequest{Password: "12345678"}
	model := models.User{}
	entity := entities.AuthUserResponse{}
	t.mockRepoV1.EXPECT().ReadUsersByEmail(t.ctx, &model).RunAndReturn(func(ctx context.Context, u *models.User) error {
		u.Auth = "$2a$06$vxnncRK0fwwjaZrEtn.ddeBccCH2BYZARKCZv.sgtKYJh8hINbfSS"
		return nil
	}).Once()

	err := t.svc.AuthUser(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.AuthUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}
	resp := entities.AuthUserResponse{
		JWT: entity.JWT,
	}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.AuthUser() = %v, want %v", entity, resp)
	}
}

func (t *ServTestSuite) Test_FindUserByEmail_FindUsersByEmail_DB_error() {
	t.wantErr = true
	req := entities.FindUserByEmailRequest{}
	t.mockRepoV1.EXPECT().FindUsersByEmail(t.ctx, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.FindUsersByEmail(t.ctx, 0, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.FindUserByEmail() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_FindUserByEmail_FindUsersByEmail_DB() {
	t.wantErr = false
	emailResp := "barko123@gmail.com"
	req := entities.FindUserByEmailRequest{
		Email: emailResp,
	}
	var model []models.User
	entity := []entities.FindUserByEmailResponse{}
	t.mockRepoV1.EXPECT().FindUsersByEmail(t.ctx, emailResp, mock.Anything, &model).RunAndReturn(func(ctx context.Context, s string, i int32, u *[]models.User) error {
		*u = append(*u, models.User{
			Email: emailResp,
		})
		return nil
	}).Once()

	err := t.svc.FindUsersByEmail(t.ctx, 0, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.CreateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

	resp := []entities.FindUserByEmailResponse{
		{
			Email: emailResp,
		},
	}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.FindUserByEmail() = %v, want %v", entity, resp)
	}
}

func (t *ServTestSuite) Test_UpdateUser_Password_LessThan8() {
	t.wantErr = true
	auth := "123"
	req := entities.UpdateUserRequest{
		Auth:        &auth,
		ConfirmPass: &auth,
	}

	err := t.svc.UpdateUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_UpdateUser_Password_MoreThan20() {
	t.wantErr = true
	auth := "12345678901234567890123"
	req := entities.UpdateUserRequest{
		Auth:        &auth,
		ConfirmPass: &auth,
	}

	err := t.svc.UpdateUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_UpdateUser_Password_And_Confirm_Not_Match() {
	t.wantErr = true
	auth1 := "123456789"
	auth2 := "12345678"
	req := entities.UpdateUserRequest{
		Auth:        &auth1,
		ConfirmPass: &auth2,
	}

	err := t.svc.UpdateUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_UpdateUser_UpdateUser_DB_Error() {
	t.wantErr = true
	auth1 := "12345678"
	auth2 := "12345678"
	req := entities.UpdateUserRequest{
		Auth:        &auth1,
		ConfirmPass: &auth2,
	}

	t.mockRepoV1.EXPECT().UpdateUsers(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.UpdateUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_UpdateUser_UpdateUser_DB_ReadUser_DB_Error() {
	t.wantErr = true
	auth1 := "12345678"
	auth2 := "12345678"
	// auth3 := "$2a$06$vxnncRK0fwwjaZrEtn.ddeBccCH2BYZARKCZv.sgtKYJh8hINbfSS"
	req := entities.UpdateUserRequest{
		Auth:        &auth1,
		ConfirmPass: &auth2,
	}

	t.mockRepoV1.EXPECT().UpdateUsers(t.ctx, mock.Anything).RunAndReturn(func(ctx context.Context, u models.User) error {
		return nil
	}).Once()

	t.mockRepoV1.EXPECT().ReadUsersById(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.UpdateUser(t.ctx, req, nil)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_UpdateUser_UpdateUser_DB_ReadUser_DB_Success() {
	t.wantErr = false
	auth1 := "12345678"
	auth2 := "12345678"
	auth3 := "$2a$06$vxnncRK0fwwjaZrEtn.ddeBccCH2BYZARKCZv.sgtKYJh8hINbfSS"
	req := entities.UpdateUserRequest{
		Auth:        &auth1,
		ConfirmPass: &auth2,
	}

	entity := entities.UpdateUserResponse{}

	t.mockRepoV1.EXPECT().UpdateUsers(t.ctx, mock.Anything).RunAndReturn(func(ctx context.Context, u models.User) error {
		return nil
	}).Once()

	t.mockRepoV1.EXPECT().ReadUsersById(t.ctx, mock.Anything).RunAndReturn(func(ctx context.Context, u *models.User) error {
		u.Auth = auth3
		return nil
	}).Once()

	err := t.svc.UpdateUser(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.UpdateUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_DeleteUser_DeleteUserById_Error() {
	t.wantErr = true
	req := entities.DeleteUserRequest{
		Id: 1,
	}

	t.mockRepoV1.EXPECT().DeleteUserById(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.DeleteUser(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.DeleteUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeleteUser_DeleteUserById_Success() {
	t.wantErr = false
	var mockId int32 = 1
	req := entities.DeleteUserRequest{
		Id: mockId,
	}

	t.mockRepoV1.EXPECT().DeleteUserById(t.ctx, mock.Anything).RunAndReturn(func(ctx context.Context, u models.User) error {
		u.DeletedAt = &t.timeGen
		return nil
	}).Once()

	err := t.svc.DeleteUser(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(nil, "Service.DeleteUser() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}
