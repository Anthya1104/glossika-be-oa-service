package data

import (
	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
)

type MockDataManager struct {
	CreateMockData
	GetMockData
}

func NewMockDataManager() *MockDataManager {
	return &MockDataManager{}
}

type CreateMockData interface {
	CreateUserInfo(userId string, userName string) (userInfo dbModel.UserInfo, err error)
}

type GetMockData interface {
	GetUserInfo(userId string) (userInfo dbModel.UserInfo, err error)
}
