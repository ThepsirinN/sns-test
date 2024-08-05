package response

import (
	"net/http"
)

const (
	CODE_SUCCESS     = 2000
	CODE_NOT_SUCCESS = 4000
	CODE_UNAUTHORIZE = 4001
	CODE_NOT_FOUND   = 4004

	MESSAGE_SUCCESS     = "success"
	MESSAGE_NOT_SUCCESS = "API failed"
	MESSAGE_UNAUTHORIZE = "unauthorize"
	MESSAGE_NOT_FOUND   = "not found"
)

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func NewResponse[T any](code int, message string, data T) Response[T] {
	return Response[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewSuccessResponse[T any](data T) (int, Response[T]) {
	return http.StatusOK, Response[T]{
		Code:    CODE_SUCCESS,
		Message: MESSAGE_SUCCESS,
		Data:    data,
	}
}

func NewNotSuccessResponse[T any](data T) (int, Response[T]) {
	return http.StatusOK, Response[T]{
		Code:    CODE_NOT_SUCCESS,
		Message: MESSAGE_NOT_SUCCESS,
		Data:    data,
	}
}

func NewUnAuthorizeResponse[T any](data T) (int, Response[T]) {
	return http.StatusOK, Response[T]{
		Code:    CODE_UNAUTHORIZE,
		Message: MESSAGE_UNAUTHORIZE,
		Data:    data,
	}
}

func NewNotFoundResponse[T any](data T) (int, Response[T]) {
	return http.StatusOK, Response[T]{
		Code:    CODE_NOT_FOUND,
		Message: MESSAGE_NOT_FOUND,
		Data:    data,
	}
}

func NewSuccessWithOutDataResponse() (int, Response[any]) {
	return http.StatusOK, Response[any]{
		Code:    CODE_SUCCESS,
		Message: MESSAGE_SUCCESS,
	}
}

func NewNotSuccessWithOutDataResponse() (int, Response[any]) {
	return http.StatusOK, Response[any]{
		Code:    CODE_NOT_SUCCESS,
		Message: MESSAGE_NOT_SUCCESS,
	}
}

func NewUnAuthorizeWithOutDataResponse() (int, Response[any]) {
	return http.StatusOK, Response[any]{
		Code:    CODE_UNAUTHORIZE,
		Message: MESSAGE_UNAUTHORIZE,
	}
}

func NewNotFoundWithOutDataResponse() (int, Response[any]) {
	return http.StatusOK, Response[any]{
		Code:    CODE_NOT_FOUND,
		Message: MESSAGE_NOT_FOUND,
	}
}
