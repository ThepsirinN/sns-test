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

func (t *handlerTestSuite) Test_CreateFriendRequest_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreateFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateFriendRequest_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.CreateFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateFriendRequest_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"dest_id":18}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().CreateFriendRequest(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.CreateFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_CreateFriendRequest_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"dest_id":18}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().CreateFriendRequest(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.CreateFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_GetAllFriendRequest_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(0))

	err := t.handlers.GetAllFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.GetAllFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_GetAllFriendRequest_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().GetAllFriendRequest(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.GetAllFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.GetAllFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_GetAllFriendRequest_Service_Error_No_Data() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().GetAllFriendRequest(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("not found data from user")).Once()

	err := t.handlers.GetAllFriendRequest(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.GetAllFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_GetAllFriendRequest_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().GetAllFriendRequest(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, gafrr1 entities.GetAllFriendRequestRequest, gafrr2 *[]entities.GetAllFriendRequestResponse) error {
		*gafrr2 = append(*gafrr2, entities.GetAllFriendRequestResponse{
			SourceId: 1,
			DestId:   2,
		})
		return nil
	}).Once()

	err := t.handlers.GetAllFriendRequest(c)
	var validateRespData response.Response[[]entities.GetAllFriendRequestResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.GetAllFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := []entities.GetAllFriendRequestResponse{
		{SourceId: 1, DestId: 2},
	}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.GetAllFriendRequest() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_ListFriend_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(0))

	err := t.handlers.ListFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListFriend_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().ListFriend(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.ListFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListFriend_Service_Error_No_Data() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().ListFriend(mock.Anything, mock.Anything, mock.Anything).Return(errors.New("not found data from user")).Once()

	err := t.handlers.ListFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_ListFriend_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodGet, "/", nil)
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(1))

	t.mockServiceV1.EXPECT().ListFriend(mock.Anything, mock.Anything, mock.Anything).RunAndReturn(func(ctx context.Context, lfr1 entities.ListFriendRequest, lfr2 *[]entities.ListFriendResponse) error {
		*lfr2 = append(*lfr2, entities.ListFriendResponse{
			FriendUserID: 2,
		})
		return nil
	}).Once()

	err := t.handlers.ListFriend(c)
	var validateRespData response.Response[[]entities.ListFriendResponse]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.ListFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}

	resp := []entities.ListFriendResponse{
		{FriendUserID: 2},
	}

	if !t.Equal(validateRespData.Data, resp) {
		t.Errorf(nil, "Handler.ListFriend() = %v, want %v", validateRespData.Data, resp)
	}
}

func (t *handlerTestSuite) Test_UpdateFriendRequestStatus_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.UpdateFriendRequestStatus(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateFriendRequestStatus() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateFriendRequestStatus_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.UpdateFriendRequestStatus(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateFriendRequestStatus() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateFriendRequestStatus_User_Not_Match_DestId() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(4))

	err := t.handlers.UpdateFriendRequestStatus(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateFriendRequestStatus() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateFriendRequestStatus_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(3))

	t.mockServiceV1.EXPECT().UpdateFriendRequestStatus(mock.Anything, mock.Anything).Return(errors.New(""))

	err := t.handlers.UpdateFriendRequestStatus(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateFriendRequestStatus() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_UpdateFriendRequestStatus_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodPatch, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(3))

	t.mockServiceV1.EXPECT().UpdateFriendRequestStatus(mock.Anything, mock.Anything).Return(nil)

	err := t.handlers.UpdateFriendRequestStatus(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.UpdateFriendRequestStatus() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteFriend_Binding_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader("{"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeleteFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.CreateFriendRequest() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteFriend_Validate_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader("{}"))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")

	err := t.handlers.DeleteFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteFriend_User_Not_In_Dest_Or_Source() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(32))

	err := t.handlers.DeleteFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteFriend_Service_Error() {
	t.wantErr = true
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(3))

	t.mockServiceV1.EXPECT().DeleteFriend(mock.Anything, mock.Anything).Return(errors.New("")).Once()

	err := t.handlers.DeleteFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}

func (t *handlerTestSuite) Test_DeleteFriend_Success() {
	t.wantErr = false
	reqValidate := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(`{"id":1,"source_id":2,"dest_id":3}`))
	reqValidate.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	reqValidate.Header.Add("Authorization", mock.Anything)
	rec := httptest.NewRecorder()
	c := t.e.NewContext(reqValidate, rec)
	c.SetPath("/")
	c.Set("user_id", int32(3))

	t.mockServiceV1.EXPECT().DeleteFriend(mock.Anything, mock.Anything).Return(nil).Once()

	err := t.handlers.DeleteFriend(c)
	var validateRespData response.Response[any]
	_ = json.Unmarshal(rec.Body.Bytes(), &validateRespData)
	if (validateRespData.Code != response.CODE_SUCCESS) != t.wantErr {
		t.Errorf(err, "Handler.DeleteFriend() error = %v,Code = %v, wantErr %v", err.Error(), validateRespData.Code, t.wantErr)
		return
	}
}
