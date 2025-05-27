package test

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/Anthya1104/gin-base-service/internal/app/model"
	"github.com/Anthya1104/gin-base-service/pkg/errcode"
	"github.com/Anthya1104/gin-base-service/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GetUserInfoTestSuite struct {
	suite.Suite
	ctx       context.Context
	targetApi string
}

// SetupSuite runs once before the suite starts
func (suite *GetUserInfoTestSuite) SetupSuite() {
	suite.targetApi = "/api/users"
	suite.ctx = context.Background()
	Setup()
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
	log.L.Info("Running Test_200_anyway")

	uri := suite.targetApi + "?userId=12345"
	w, _ := HttpGet(uri, nil)
	resp := model.GetUserInfoResp{}
	json.Unmarshal(w.Body.Bytes(), &resp)

	t := suite.T()
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, getRespVersion(w))
	assert.Equal(t, "12345", resp.Data.UserID)
}
