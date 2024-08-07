package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sns-barko/utility/response"
	"sns-barko/v1/entities"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func (t *handlerTestSuite) Test_CreateUser_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateUser_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"x"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateUser_Password_MissMatch() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john.doe@example.com","first_name":"John","last_name":"Doe","password":"password123","confirm_password":"password124"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateUser_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john.doe@example.com","first_name":"John","last_name":"Doe","password":"password123","confirm_password":"password123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().CreateUser(ctx, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.CreateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateUser_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john.doe@example.com","first_name":"John","last_name":"Doe","password":"password123","confirm_password":"password123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().CreateUser(ctx, mock.Anything).Return(nil).Once()

	err := t.handlers.CreateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AuthUser_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.AuthUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AuthUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AuthUser_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"x"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.AuthUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AuthUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AuthUser_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john.doe@example.com","password":"password123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().AuthUser(ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.AuthUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AuthUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AuthUser_Service_Succes() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"email":"john.doe@example.com","password":"password123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	ctx := c.Request().Context()
	mockJWT := "mockJWT"

	t.mockServiceV1.EXPECT().AuthUser(ctx, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, aur1 entities.AuthUserRequest, aur2 *entities.AuthUserResponse) error {
		aur2.JWT = mockJWT
		return nil
	}).Once()

	err := t.handlers.AuthUser(c)
	var validateRespData response.Response[entities.AuthUserResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AuthUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := entities.AuthUserResponse{
		JWT: mockJWT,
	}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.AuthUser() = %v, want %v", validateRespData.Data.JWT, resp)
	}
}

func (t *handlerTestSuite) Test_FindUserByEmail_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.FindUserByEmail(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.FindUserByEmail() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_FindUserByEmail_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.SetParamNames("email")
	c.SetParamValues("barko123@gmail.com")
	c.Set("user_id", 0)

	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().FindUsersByEmail(ctx, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.FindUserByEmail(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.FindUserByEmail() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_FindUserByEmail_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.SetParamNames("email")
	c.SetParamValues("barko123@gmail.com")
	c.Set("user_id", 0)

	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().FindUsersByEmail(ctx, mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, i int32, fuber1 entities.FindUserByEmailRequest, fuber2 *[]entities.FindUserByEmailResponse) error {

		*fuber2 = append(*fuber2, entities.FindUserByEmailResponse{
			Id:    0,
			Email: "barko123@gmail.com",
		},
		)
		return nil
	}).Once()

	err := t.handlers.FindUserByEmail(c)
	var validateRespData response.Response[[]entities.FindUserByEmailResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.FindUserByEmail() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := []entities.FindUserByEmailResponse{
		{Id: 0, Email: "barko123@gmail.com"},
	}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.FindUserByEmail() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_UpdateUser_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.UpdateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateUser_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"first_name":"123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", 0)
	err := t.handlers.UpdateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateUser_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"first_name":"123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))
	ctx := c.Request().Context()

	t.mockServiceV1.EXPECT().UpdateUser(ctx, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.UpdateUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateUser_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"first_name":"123"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))
	ctx := c.Request().Context()
	mockJWT := "mock"

	t.mockServiceV1.EXPECT().UpdateUser(ctx, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, uur1 entities.UpdateUserRequest, uur2 *entities.UpdateUserResponse) error {
		uur2.JWT = mockJWT
		return nil
	}).Once()

	err := t.handlers.UpdateUser(c)
	var validateRespData response.Response[entities.UpdateUserResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := entities.UpdateUserResponse{
		JWT: mockJWT,
	}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.UpdateUser() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_DeleteUser_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(0))

	err := t.handlers.DeleteUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteUser_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().DeleteUser(mock.Anything, mock.Anything).Return(errors.New(""))

	err := t.handlers.DeleteUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteUser_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().DeleteUser(mock.Anything, mock.Anything).Return(nil)

	err := t.handlers.DeleteUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}
