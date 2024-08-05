package handlers

import (
	"errors"
	"sns-barko/utility/logger"
	"sns-barko/utility/response"
	"sns-barko/v1/entities"
	"strings"

	"github.com/labstack/echo/v4"
)

func (h *HandlersV1) CreateFriendRequest(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.CreateFriendRequestRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.SourceId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.CreateFriendRequest(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) GetAllFriendRequest(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.GetAllFriendRequestRequest

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.UserId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	var resp []entities.GetAllFriendRequestResponse
	if err := h.service.GetAllFriendRequest(ctx, req, &resp); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		if strings.Contains(err.Error(), "not found data from user") {
			errRespCode, errResp := response.NewNotFoundWithOutDataResponse()
			return c.JSON(errRespCode, errResp)
		}
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessResponse(resp)

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) ListFriend(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.ListFriendRequest

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.UserId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	var resp []entities.ListFriendResponse
	if err := h.service.ListFriend(ctx, req, &resp); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		if strings.Contains(err.Error(), "not found data from user") {
			errRespCode, errResp := response.NewNotFoundWithOutDataResponse()
			return c.JSON(errRespCode, errResp)
		}

		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessResponse(resp)

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) UpdateFriendRequestStatus(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.UpdateFriendRequestStatusRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)

	if dataInt32 != req.DestId {
		err := errors.New("don't have permission to update")
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.UpdateFriendRequestStatus(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) DeleteFriend(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.DeleteFriendRequestRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)

	if dataInt32 != req.DestId && dataInt32 != req.SourceId {
		err := errors.New("don't have permission to delete")
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.DeleteFriend(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}
