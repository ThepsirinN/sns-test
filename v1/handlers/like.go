package handlers

import (
	"sns-barko/utility/logger"
	"sns-barko/utility/response"
	"sns-barko/v1/entities"

	"github.com/labstack/echo/v4"
)

func (h *HandlersV1) AddLike(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.AddLikeRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}
	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.UserId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.AddLike(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) DeleteLike(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.DeleteLikeRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}
	data := c.Get("user_id")
	dataInt32, _ := data.(int32)
	req.UserId = dataInt32

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.DeleteLike(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}
