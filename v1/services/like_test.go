package services

import (
	"errors"
	"sns-barko/v1/entities"

	"github.com/stretchr/testify/mock"
)

func (t *ServTestSuite) Test_AddLike_DB_Error() {
	t.wantErr = true
	req := entities.AddLikeRequest{}

	t.mockRepoV1.EXPECT().UpdateLike(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.AddLike(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.AddLike() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_AddLike_DB_Success() {
	t.wantErr = false
	req := entities.AddLikeRequest{}

	t.mockRepoV1.EXPECT().UpdateLike(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	err := t.svc.AddLike(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.AddLike() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_DeleteLike_DB_Error() {
	t.wantErr = true
	req := entities.DeleteLikeRequest{}

	t.mockRepoV1.EXPECT().DeleteLike(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.DeleteLike(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeleteLike() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeleteLike_DB_Success() {
	t.wantErr = false
	req := entities.DeleteLikeRequest{}

	t.mockRepoV1.EXPECT().DeleteLike(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	err := t.svc.DeleteLike(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeleteLike() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}
