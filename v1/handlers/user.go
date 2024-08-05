package handlers

import (
	"errors"
	"sns-barko/utility/logger"
	"sns-barko/utility/response"
	"sns-barko/v1/entities"

	"github.com/labstack/echo/v4"
)

func (h *HandlersV1) CreateUser(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.CreateUserRequest

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

	if req.Auth != req.ConfirmPass {
		err := errors.New("password and confirm password miss match")
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	if err := h.service.CreateUser(ctx, req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) AuthUser(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.AuthUserRequest

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

	resp := new(entities.AuthUserResponse)
	err := h.service.AuthUser(ctx, req, resp)
	if err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessResponse(resp)

	return c.JSON(successRespCode, successResp)

}

func (h *HandlersV1) FindUserByEmail(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.FindUserByEmailRequest

	req.Email = c.Param("email")
	data := c.Get("user_id")
	userId, _ := data.(int32)

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	var resp []entities.FindUserByEmailResponse
	err := h.service.FindUsersByEmail(ctx, userId, req, &resp)
	if err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessResponse(resp)

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) UpdateUser(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	data := c.Get("user_id")
	req.Id, _ = data.(int32)

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}
	resp := new(entities.UpdateUserResponse)
	err := h.service.UpdateUser(ctx, req, resp)
	if err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessResponse(resp)

	return c.JSON(successRespCode, successResp)
}

func (h *HandlersV1) DeleteUser(c echo.Context) error {
	echoReq := c.Request()
	ctx := echoReq.Context()
	var req entities.DeleteUserRequest

	data := c.Get("user_id")
	req.Id, _ = data.(int32)

	if err := c.Validate(req); err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	err := h.service.DeleteUser(ctx, req)
	if err != nil {
		logger.Error(ctx, err)
		errRespCode, errResp := response.NewNotSuccessResponse(err.Error())
		return c.JSON(errRespCode, errResp)
	}

	successRespCode, successResp := response.NewSuccessWithOutDataResponse()

	return c.JSON(successRespCode, successResp)

}
