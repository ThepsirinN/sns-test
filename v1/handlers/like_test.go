package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sns-barko/utility/response"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

func (t *handlerTestSuite) Test_AddLike_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.AddLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AddLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AddLike_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.AddLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AddLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AddLike_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"post_id":12}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().AddLike(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.AddLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AddLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_AddLike_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"post_id":12}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().AddLike(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.AddLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.AddLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteLike_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeleteLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteLike_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeleteLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteLike_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"post_id":12}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().DeleteLike(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.DeleteLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteLike_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"post_id":12}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().DeleteLike(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.DeleteLike(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteLike() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}
