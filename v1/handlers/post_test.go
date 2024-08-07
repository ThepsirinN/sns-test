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

func (t *handlerTestSuite) Test_CreatePost_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreatePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreatePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreatePost_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreatePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreatePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreatePost_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"post_data":"asdf"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().CreatePost(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.CreatePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreatePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreatePost_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"post_data":"asdf"}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().CreatePost(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.CreatePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreatePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListAllPostFromUser_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.ListAllPostFromUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListAllPostFromUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListAllPostFromUser_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ListAllPostFromUser(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.ListAllPostFromUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListAllPostFromUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListAllPostFromUser_Service_Error_No_Data() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ListAllPostFromUser(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("not found data from user")).Once()

	err := t.handlers.ListAllPostFromUser(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListAllPostFromUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListAllPostFromUser_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ListAllPostFromUser(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, lapfur1 entities.ListAllPostFromUserRequest, lapfur2 *[]entities.ListAllPostFromUserResponse) error {
		*lapfur2 = append(*lapfur2, entities.ListAllPostFromUserResponse{OwnerId: int32(2)})
		return nil
	}).Once()

	err := t.handlers.ListAllPostFromUser(c)
	var validateRespData response.Response[[]entities.ListAllPostFromUserResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListAllPostFromUser() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := []entities.ListAllPostFromUserResponse{{OwnerId: int32(2)}}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.ListAllPostFromUser() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_ReadPostByPostId_Biding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetParamNames("id")
	c.SetParamValues("{")
	c.SetPath("/")

	err := t.handlers.ReadPostByPostId(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ReadPostByPostId() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ReadPostByPostId_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.ReadPostByPostId(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ReadPostByPostId() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ReadPostByPostId_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.SetParamNames("id")
	c.SetParamValues("3")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ReadPostByPostId(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.ReadPostByPostId(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ReadPostByPostId() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ReadPostByPostId_Service_Error_No_Data() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.SetParamNames("id")
	c.SetParamValues("3")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ReadPostByPostId(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("not found data from user")).Once()

	err := t.handlers.ReadPostByPostId(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ReadPostByPostId() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ReadPostByPostId_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.SetParamNames("id")
	c.SetParamValues("3")
	c.Set("user_id", int32(2))

	t.mockServiceV1.EXPECT().ReadPostByPostId(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, rpbpir1 entities.ReadPostByPostIdRequest, rpbpir2 *entities.ReadPostByPostIdResponse) error {
		rpbpir2.OwnerId = int32(2)
		rpbpir2.Id = int32(3)
		return nil
	}).Once()

	err := t.handlers.ReadPostByPostId(c)
	var validateRespData response.Response[entities.ReadPostByPostIdResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ReadPostByPostId() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := entities.ReadPostByPostIdResponse{Id: int32(3), OwnerId: int32(2)}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.ReadPostByPostId() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_UpdatePostData_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.UpdatePostData(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdatePostData() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdatePostData_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.UpdatePostData(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdatePostData() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdatePostData_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":213}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().UpdatePostData(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.UpdatePostData(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdatePostData() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdatePostData_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":213}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().UpdatePostData(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.UpdatePostData(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdatePostData() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeletePost_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeletePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeletePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeletePost_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeletePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeletePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeletePost_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":213}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().DeletePost(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.DeletePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeletePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeletePost_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":213}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(12))

	t.mockServiceV1.EXPECT().DeletePost(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.DeletePost(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeletePost() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}
