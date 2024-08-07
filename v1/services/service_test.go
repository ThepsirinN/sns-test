package services

import (
	"context"
	"sns-barko/config"
	"sns-barko/v1/services/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ServTestSuite struct {
	suite.Suite

	mockRepoV1  *mocks.RepoV1Interface
	mockCacheV1 *mocks.CacheV1
	wantErr     bool

	svc *serviceV1

	timeGen time.Time

	ctx context.Context
}

func (t *ServTestSuite) SetupTest() {
	t.mockRepoV1 = mocks.NewRepoV1Interface(t.T())
	t.mockCacheV1 = mocks.NewCacheV1(t.T())

	mockAppSecret := &config.Secret{}
	mockAppSecret.User.Hash.Secret = "secret"
	mockAppSecret.User.Hash.Cost = 6
	mockAppSecret.User.JWT.Secret = "secret"
	mockAppSecret.User.JWT.ExpDuration = 15 * time.Second

	t.ctx = context.Background()
	t.mockRepoV1.EXPECT().AutoMigrate(t.ctx).Return(nil).Once()

	t.timeGen = time.Date(2024, 8, 1, 23, 59, 59, 0, time.Local)
	t.svc = New(t.ctx, t.mockRepoV1, t.mockCacheV1, mockAppSecret)

	// t.svc.customizeNowFunc(func() time.Time {
	//  return mockNow
	// })
}

// func (t *ServTestSuite) TearDownTest() {
// }

func Test_Run(t *testing.T) {
	suite.Run(t, new(ServTestSuite))
}
