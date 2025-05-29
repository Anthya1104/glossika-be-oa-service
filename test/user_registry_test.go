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

type UserRegistryTestSuite struct {
	suite.Suite
	ctx            context.Context
	testContainers container.TestContainers
	targetApi      string
}

// SetupSuite runs once before the suite starts
func (suite *UserRegistryTestSuite) SetupSuite() {
	suite.targetApi = "/api/v1/users"
	suite.ctx = context.Background()
	suite.testContainers = container.NewTestContainers(suite.ctx)
	Setup()
}

// TearDownSuite runs once after all tests in the suite have completed
func (suite *UserRegistryTestSuite) TearDownSuite() {
	database.GetSqlDb().CloseConnection()
	suite.testContainers.TearDownContainers(suite.ctx)
}

// SetupTest runs before each test in the suite
func (suite *UserRegistryTestSuite) SetupTest() {
}

// TearDownTest runs after each test in the suite
func (suite *UserRegistryTestSuite) TearDownTest() {
}

func TestUserRegistryTestSuite(t *testing.T) {
	log.L.Info("----- RUN TestUserRegistryTestSuite -----")
	suite.Run(t, new(UserRegistryTestSuite))
	log.L.Info("----- FINISH TestUserRegistryTestSuite -----")
}

func (suite *UserRegistryTestSuite) Test_400_when_missing_user_email() {
	log.L.Info("Running Test_400_when_missing_user_email")

	reqBody := model.UserRegisterReq{Password: "fakePass#"}
	b, _ := json.Marshal(reqBody)
	w, _ := HttpPost(suite.targetApi, string(b), nil)

	t := suite.T()
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
	assert.Equal(t, string(errcode.BadRequestBody), getRespCustomErrorcode(w))
	assert.Equal(t, errcode.ErrCodeMsg[errcode.BadRequestBody], getRespCustomErrorMsg(w))
}

func (suite *UserRegistryTestSuite) Test_400_when_registered_user_email_exists() {
	log.L.Info("Running Test_400_when_registered_user_email_exists")
	resetDbForTesting()

	// Create mock data
	mockDataManager := data.NewMockDataManager()
	mockDataManager.CreateUser("test@email.com")

	reqBody := model.UserRegisterReq{Email: "test@email.com", Password: "fakePass#"}
	b, _ := json.Marshal(reqBody)
	w, _ := HttpPost(suite.targetApi, string(b), nil)

	t := suite.T()
	assert.Equal(t, http.StatusConflict, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
	assert.Equal(t, string(errcode.DBDuplicatedUser), getRespCustomErrorcode(w))
}

func (suite *UserRegistryTestSuite) Test_200_when_user_registry_succeed() {
	resetDbForTesting()
	log.L.Info("Running Test_200_when_user_registry_succeed")

	reqBody := model.UserRegisterReq{Email: "test@email.com", Password: "fakePass#"}
	b, _ := json.Marshal(reqBody)
	w, _ := HttpPost(suite.targetApi, string(b), nil)

	resp := model.GetUserInfoResp{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	t := suite.T()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
}
