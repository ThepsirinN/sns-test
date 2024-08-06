package handlers

import (
	"sns-barko/utility/logger"
	"sns-barko/utility/response"
	"sns-barko/v1/entities"
	"strings"

	"github.com/labstack/echo/v4"
)

// func ()
func (h *HandlersV1) CreatePost(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.CreatePostRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.OwnerId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.CreatePost(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) ListAllPostFromUser(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.ListAllPostFromUserRequest

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.OwnerId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	var resp []entities.ListAllPostFromUserResponse
	if err := h.service.ListAllPostFromUser(ctx, req, &resp); err != nil {
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

func (h *HandlersV1) ReadPostByPostId(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.ReadPostByPostIdRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.OwnerId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	var resp entities.ReadPostByPostIdResponse
	if err := h.service.ReadPostByPostId(ctx, req, &resp); err != nil {
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

func (h *HandlersV1) UpdatePostData(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.UpdatePostRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}
	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.OwnerId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.UpdatePostData(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) DeletePost(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.DeletePostRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}
	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.OwnerId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.DeletePost(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}
