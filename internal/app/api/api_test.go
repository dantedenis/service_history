package api

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	mock_contract "service_history/internal/app/contract/mock"
	"service_history/pkg/config"
	mock_config "service_history/pkg/config/mock"
	"testing"
)

func TestNewApiServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	configMock := mock_config.NewMockIConfig(ctrl)
	configMock.EXPECT().GetPort().Return("")
	configMock.EXPECT().GetSQL().Return(&config.SQL{})
	configMock.EXPECT().GetPeriod().Return(0)

	apiServ := NewApiServer(configMock)
	assert.NotNil(t, apiServ)
}

func TestServer_StartServ(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_contract.NewMockIProvider(ctrl)
	requestMock := mock_contract.NewMockIRequester(ctrl)
	serviceMock := mock_contract.NewMockIService(ctrl)

	repoMock.EXPECT().Open().Return(nil)
	repoMock.EXPECT().GetConn().Return(&sql.DB{}).AnyTimes()

	requestMock.EXPECT().Start(repoMock.GetConn()).Return(nil)

	serviceMock.EXPECT().Run(context.Background()).Return(nil)

	apiServ := &Server{}
	apiServ.repo = repoMock
	apiServ.requester = requestMock
	apiServ.service = serviceMock

	assert.Nil(t, apiServ.StartServ())
}

func TestServer_StartServ_err(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_contract.NewMockIProvider(ctrl)
	requestMock := mock_contract.NewMockIRequester(ctrl)
	serviceMock := mock_contract.NewMockIService(ctrl)

	repoMock.EXPECT().Open().Return(nil).MinTimes(1)
	repoMock.EXPECT().GetConn().Return(&sql.DB{}).AnyTimes()

	requestMock.EXPECT().Start(repoMock.GetConn()).Return(errors.New("")).MinTimes(1)

	apiServ := &Server{}
	apiServ.repo = repoMock
	apiServ.requester = requestMock
	apiServ.service = serviceMock

	assert.NotNil(t, apiServ.StartServ())

	assert.Nil(t, os.Setenv("RPC_PORT", ""))
	assert.NotNil(t, apiServ.StartServ())
}

func TestServer_StartServ_err2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_contract.NewMockIProvider(ctrl)
	requestMock := mock_contract.NewMockIRequester(ctrl)
	serviceMock := mock_contract.NewMockIService(ctrl)

	repoMock.EXPECT().Open().Return(errors.New("")).MinTimes(2)

	apiServ := &Server{}
	apiServ.repo = repoMock
	apiServ.requester = requestMock
	apiServ.service = serviceMock

	assert.NotNil(t, apiServ.StartServ())
}
