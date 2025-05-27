package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/model"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/Anthya1104/glossika-be-oa-service/test/container"
	"github.com/Anthya1104/glossika-be-oa-service/test/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetUserInfoTestSuite struct {
	suite.Suite
	ctx            context.Context
	testContainers container.TestContainers
	targetApi      string
}

// SetupSuite runs once before the suite starts
func (suite *GetUserInfoTestSuite) SetupSuite() {
	suite.targetApi = "/api/users"
	suite.ctx = context.Background()
	suite.testContainers = container.NewTestContainers(suite.ctx)
	Setup()
}

// TearDownSuite runs once after all tests in the suite have completed
func (suite *GetUserInfoTestSuite) TearDownSuite() {
	database.GetSqlDb().CloseConnection()
	suite.testContainers.TearDownContainers(suite.ctx)
}

// SetupTest runs before each test in the suite
func (suite *GetUserInfoTestSuite) SetupTest() {
}

// TearDownTest runs after each test in the suite
func (suite *GetUserInfoTestSuite) TearDownTest() {
}

func TestGetUserInfoTestSuite(t *testing.T) {
	log.L.Info("----- RUN TestGetUserInfoTestSuite -----")
	suite.Run(t, new(GetUserInfoTestSuite))
	log.L.Info("----- FINISH TestGetUserInfoTestSuite -----")
}

func (suite *GetUserInfoTestSuite) Test_400_when_header_missing_user_id() {
	log.L.Info("Running Test_400_when_header_missing_user_id")

	w, _ := HttpGet(suite.targetApi, nil)

	t := suite.T()
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
	assert.Equal(t, string(errcode.BadQuery), getRespCustomErrorcode(w))
	assert.Equal(t, errcode.ErrCodeMsg[errcode.BadQuery], getRespCustomErrorMsg(w))
}

func (suite *GetUserInfoTestSuite) Test_200_anyway() {
	resetDbForTesting()
	log.L.Info("Running Test_200_anyway")

	// Create mock data
	mockDataManager := data.NewMockDataManager()
	mockDataManager.CreateUserInfo("12345", "test_user")

	uri := suite.targetApi + "?userId=12345"
	w, _ := HttpGet(uri, nil)
	resp := model.GetUserInfoResp{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	t := suite.T()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
	assert.Equal(t, "12345", resp.Data.UserID)
	assert.Equal(t, "test_user", resp.Data.UserName)
}
