package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sns-barko/constant"
	"sns-barko/v1/entities"

	"github.com/stretchr/testify/mock"
)

func (t *ServTestSuite) Test_CreateFriend_DB_Error() {
	t.wantErr = true
	req := entities.CreateFriendRequestRequest{}

	t.mockRepoV1.EXPECT().CreateFriendRequest(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.CreateFriendRequest(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.CreateFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_CreateFriend_DB_Cache_Error() {
	t.wantErr = true
	req := entities.CreateFriendRequestRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().CreateFriendRequest(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(errors.New("")).Once()

	err := t.svc.CreateFriendRequest(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.CreateFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_CreateFriend_DB_Cache_Success() {
	t.wantErr = false
	req := entities.CreateFriendRequestRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().CreateFriendRequest(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(nil).Once()

	err := t.svc.CreateFriendRequest(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.CreateFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_GetAllFriendRequest_Data_From_Cache() {
	t.wantErr = false
	req := entities.GetAllFriendRequestRequest{}
	entity := []entities.GetAllFriendRequestResponse{}

	mockCache := []entities.GetAllFriendRequestResponse{}
	mockCacheByte, _ := json.Marshal(mockCache)
	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(mockCacheByte, nil).Once()

	err := t.svc.GetAllFriendRequest(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.GetAllFriendRequest() error = %v, wantErr %v", err, t.wantErr)
		return
	}

	resp := []entities.GetAllFriendRequestResponse{}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.GetAllFriendRequest() = %v, want %v", entity, resp)
	}

}

func (t *ServTestSuite) Test_GetAllFriendRequest_Data_From_Cache_Error_DB_Error() {
	t.wantErr = true
	req := entities.GetAllFriendRequestRequest{}
	entity := []entities.GetAllFriendRequestResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().GetAllFriendRequest(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.GetAllFriendRequest(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.GetAllFriendRequest() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_GetAllFriendRequest_Data_From_Cache_Marshal_Error_DB_Error() {
	t.wantErr = true
	req := entities.GetAllFriendRequestRequest{}
	entity := []entities.GetAllFriendRequestResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return([]byte{}, nil).Once()
	t.mockRepoV1.EXPECT().GetAllFriendRequest(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.GetAllFriendRequest(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.GetAllFriendRequest() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_GetAllFriendRequest_Data_From_Cache_Error_DB_SetCache_Error() {
	t.wantErr = true
	req := entities.GetAllFriendRequestRequest{}
	entity := []entities.GetAllFriendRequestResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().GetAllFriendRequest(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().SetDataToCache(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.GetAllFriendRequest(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.GetAllFriendRequest() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_GetAllFriendRequest_Data_From_Cache_Error_DB_SetCache_Success() {
	t.wantErr = false
	req := entities.GetAllFriendRequestRequest{UserId: 1}
	entity := []entities.GetAllFriendRequestResponse{}

	mockRespDB := []entities.GetAllFriendRequestResponse{
		{Id: 1, SourceId: 1, Status: "1"},
		{Id: 2, DestId: 1, Status: "2"},
	}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().GetAllFriendRequest(t.ctx, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, i int32, gafrr *[]entities.GetAllFriendRequestResponse) error {
		*gafrr = mockRespDB
		return nil
	}).Once()
	t.mockCacheV1.EXPECT().SetDataToCache(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	err := t.svc.GetAllFriendRequest(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.GetAllFriendRequest() error = %v, wantErr %v", err, t.wantErr)
		return
	}
	resp := []entities.GetAllFriendRequestResponse{
		{Id: 1, SourceId: 1, Status: constant.FRIEND_STATUS_PENDING_APPROVE_MESSAGE},
		{Id: 2, DestId: 1, Status: constant.FRIEND_STATUS_PENDING_WAITING_APPROVE_MESSAGE},
	}
	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.GetAllFriendRequest() = %v, want %v", entity, resp)
	}

}

func (t *ServTestSuite) Test_ListFriend_Data_From_Cache() {
	t.wantErr = false
	req := entities.ListFriendRequest{}
	entity := []entities.ListFriendResponse{}

	mockCache := []entities.ListFriendResponse{}
	mockCacheByte, _ := json.Marshal(mockCache)
	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(mockCacheByte, nil).Once()

	err := t.svc.ListFriend(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}

	resp := []entities.ListFriendResponse{}

	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.ListFriend() = %v, want %v", entity, resp)
	}

}

func (t *ServTestSuite) Test_ListFriend_Data_From_Cache_Error_DB_Error() {
	t.wantErr = true
	req := entities.ListFriendRequest{}
	entity := []entities.ListFriendResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().ListFriend(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.ListFriend(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_ListFriend_Data_From_Cache_Marshal_Error_DB_Error() {
	t.wantErr = true
	req := entities.ListFriendRequest{}
	entity := []entities.ListFriendResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return([]byte{}, nil).Once()
	t.mockRepoV1.EXPECT().ListFriend(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.ListFriend(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_ListFriend_Data_From_Cache_Error_DB_SetCache_Error() {
	t.wantErr = true
	req := entities.ListFriendRequest{}
	entity := []entities.ListFriendResponse{}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().ListFriend(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().SetDataToCache(t.ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.ListFriend(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}

}

func (t *ServTestSuite) Test_ListFriend_Data_From_Cache_Error_DB_SetCache_Success() {
	t.wantErr = false
	req := entities.ListFriendRequest{UserId: 1}
	entity := []entities.ListFriendResponse{}

	mockRespDB := []entities.ListFriendQuery{
		{SourceId: 1, DestId: 3},
		{SourceId: 2, DestId: 1},
	}

	t.mockCacheV1.EXPECT().GetDataFromCache(t.ctx, mock.Anything).Return(nil, errors.New("")).Once()
	t.mockRepoV1.EXPECT().ListFriend(t.ctx, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, i int32, lfq *[]entities.ListFriendQuery) error {
		*lfq = mockRespDB
		return nil
	}).Once()
	t.mockCacheV1.EXPECT().SetDataToCache(t.ctx, mock.Anything, mock.Anything).Return(nil).Once()

	err := t.svc.ListFriend(t.ctx, req, &entity)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.ListFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
	resp := []entities.ListFriendResponse{
		{FriendUserID: 3},
		{FriendUserID: 2},
	}
	if !t.Equal(entity, resp) {
		t.Errorf(nil, "Service.GetAllFriendRequest() = %v, want %v", entity, resp)
	}

}

func (t *ServTestSuite) Test_UpdateFriendRequest_DB_Error() {
	t.wantErr = true
	req := entities.UpdateFriendRequestStatusRequest{}

	t.mockRepoV1.EXPECT().UpdateFriendRequestStatus(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.UpdateFriendRequestStatus(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.UpdateFriendRequestStatus() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_UpdateFriendRequest_DB_Cache_Error() {
	t.wantErr = true
	req := entities.UpdateFriendRequestStatusRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().UpdateFriendRequestStatus(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(errors.New("")).Once()

	err := t.svc.UpdateFriendRequestStatus(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.UpdateFriendRequestStatus() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_UpdateFriendRequest_DB_Cache_Success() {
	t.wantErr = false
	req := entities.UpdateFriendRequestStatusRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().UpdateFriendRequestStatus(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(nil).Once()

	err := t.svc.UpdateFriendRequestStatus(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.UpdateFriendRequestStatus() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeleteFriendRequest_DB_Error() {
	t.wantErr = true
	req := entities.DeleteFriendRequestRequest{}

	t.mockRepoV1.EXPECT().DeleteFriend(t.ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.svc.DeleteFriend(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeleteFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeleteFriendRequest_DB_Cache_Error() {
	t.wantErr = true
	req := entities.DeleteFriendRequestRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().DeleteFriend(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(errors.New(""))

	err := t.svc.DeleteFriend(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeleteFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}

func (t *ServTestSuite) Test_DeleteFriendRequest_DB_Cache_Success() {
	t.wantErr = false
	req := entities.DeleteFriendRequestRequest{}
	key := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(0),
	}

	t.mockRepoV1.EXPECT().DeleteFriend(t.ctx, mock.Anything).Return(nil).Once()
	t.mockCacheV1.EXPECT().ClearMultipleCacheKey(t.ctx, key[0], key[1], key[2], key[3]).Return(nil)

	err := t.svc.DeleteFriend(t.ctx, req)
	if (err != nil) != t.wantErr {
		t.Errorf(err, "Service.DeleteFriend() error = %v, wantErr %v", err, t.wantErr)
		return
	}
}
