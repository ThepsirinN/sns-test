package services

import (
	"context"
	"encoding/json"
	"fmt"
	"sns-barko/constant"
	"sns-barko/v1/entities"
	"sns-barko/v1/models"
)

func (s *serviceV1) CreateFriendRequest(ctx context.Context, req entities.CreateFriendRequestRequest) error {
	model := models.Friend{
		SourceId: req.SourceId,
		DestId:   req.DestId,
		Status:   constant.FRIEND_STATUS_PENDING,
	}

	err := s.repoV1.CreateFriendRequest(ctx, model)
	if err != nil {
		return err
	}

	cacheKey := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.DestId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.DestId),
	}

	err = s.cacheV1.ClearMultipleCacheKey(ctx, cacheKey...)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) GetAllFriendRequest(ctx context.Context, req entities.GetAllFriendRequestRequest, resp *[]entities.GetAllFriendRequestResponse) error {
	data, err := s.cacheV1.GetDataFromCache(ctx, constant.CACHE_FRIEND_REQUEST_KEY+fmt.Sprint(req.UserId))
	if err == nil {
		err := json.Unmarshal(data, resp)
		if err == nil {
			return nil
		}
	}

	var model []entities.GetAllFriendRequestResponse
	err = s.repoV1.GetAllFriendRequest(ctx, req.UserId, &model)
	if err != nil {
		return err
	}

	i := 0
	lengthModel := len(model)

	changeStatus := func(userId int32, sourceId int32) string {
		if userId == sourceId {
			return constant.FRIEND_STATUS_PENDING_APPROVE_MESSAGE
		}
		return constant.FRIEND_STATUS_PENDING_WAITING_APPROVE_MESSAGE
	}

	for i < lengthModel {
		model[i].Status = changeStatus(req.UserId, model[i].SourceId)
		*resp = append(*resp, model[i])
		i++
	}

	byteData, err := json.Marshal(*resp)
	if err != nil {
		return err
	}

	err = s.cacheV1.SetDataToCache(ctx, constant.CACHE_FRIEND_REQUEST_KEY+fmt.Sprint(req.UserId), byteData)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) ListFriend(ctx context.Context, req entities.ListFriendRequest, resp *[]entities.ListFriendResponse) error {
	data, err := s.cacheV1.GetDataFromCache(ctx, constant.CACHE_FRIEND_LIST_KEY+fmt.Sprint(req.UserId))
	if err == nil {
		err := json.Unmarshal(data, resp)
		if err == nil {
			return nil
		}
	}

	var model []entities.ListFriendQuery
	err = s.repoV1.ListFriend(ctx, req.UserId, &model)
	if err != nil {
		return err
	}
	i := 0
	lengthModel := len(model)

	checkFriendData := func(queryData entities.ListFriendQuery, userId int32) entities.ListFriendResponse {
		if queryData.SourceId == userId {
			return entities.ListFriendResponse{
				Id:               queryData.Id,
				FriendUserID:     queryData.DestId,
				FriendEmail:      queryData.DestEmail,
				FriendFirstName:  queryData.DestFirstName,
				FriendLastName:   queryData.DestLastName,
				FriendProfileImg: queryData.DestProfileImg,
			}
		}
		return entities.ListFriendResponse{
			Id:               queryData.Id,
			FriendUserID:     queryData.SourceId,
			FriendEmail:      queryData.SourceEmail,
			FriendFirstName:  queryData.SourceFirstName,
			FriendLastName:   queryData.SourceLastName,
			FriendProfileImg: queryData.SourceProfileImg,
		}
	}

	for i < lengthModel {
		data := checkFriendData(model[i], req.UserId)
		*resp = append(*resp, data)
		i++
	}

	byteData, err := json.Marshal(*resp)
	if err != nil {
		return err
	}

	err = s.cacheV1.SetDataToCache(ctx, constant.CACHE_FRIEND_LIST_KEY+fmt.Sprint(req.UserId), byteData)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) UpdateFriendRequestStatus(ctx context.Context, req entities.UpdateFriendRequestStatusRequest) error {
	model := models.Friend{
		Id:     req.Id,
		Status: constant.FRIEND_STATUS_SUCCESS,
	}

	err := s.repoV1.UpdateFriendRequestStatus(ctx, model)
	if err != nil {
		return err
	}

	cacheKey := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.DestId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.DestId),
	}

	err = s.cacheV1.ClearMultipleCacheKey(ctx, cacheKey...)
	if err != nil {
		return err
	}

	return nil
}

func (s *serviceV1) DeleteFriend(ctx context.Context, req entities.DeleteFriendRequestRequest) error {
	model := models.Friend{
		Id:       req.Id,
		SourceId: req.SourceId,
		DestId:   req.DestId,
	}

	err := s.repoV1.DeleteFriend(ctx, model)
	if err != nil {
		return err
	}

	cacheKey := []string{
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_REQUEST_KEY + fmt.Sprint(req.DestId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.SourceId),
		constant.CACHE_FRIEND_LIST_KEY + fmt.Sprint(req.DestId),
	}

	err = s.cacheV1.ClearMultipleCacheKey(ctx, cacheKey...)
	if err != nil {
		return err
	}

	return nil
}
