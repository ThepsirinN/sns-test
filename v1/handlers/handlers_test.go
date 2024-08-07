package handlers

import (
	"sns-barko/config"
	"sns-barko/constant"
	"sns-barko/utility/logger"
	"sns-barko/v1/handlers/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		// return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}

type handlerTestSuite struct {
	suite.Suite
	e             *echo.Echo
	mockServiceV1 *mocks.ServicesV1Interface
	handlers      *HandlersV1
	wantErr       bool
}

func (t *handlerTestSuite) SetupTest() {
	var cfg config.Config
	cfg.Log.Env = constant.ENV_LOCAL
	logger.InitLogger(&cfg)
	defer logger.Sync()
	t.mockServiceV1 = mocks.NewServicesV1Interface(t.T())
	t.e = echo.New()
	t.e.Validator = &customValidator{validator: validator.New()}
	t.handlers = New(t.mockServiceV1)

}

func Test_Run(t *testing.T) {
	suite.Run(t, new(handlerTestSuite))
}
