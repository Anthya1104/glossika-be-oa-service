package data

import (
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
)

func (m *MockDataManager) CreateUser(email string) (user dbModel.User, err error) {
	user = dbModel.User{
		Email:    email,
		Password: "fake-pass-word",
	}
	if err = database.GetSqlDb().Orm.Create(&user).Error; err != nil {
		return
	}

	return
}

func (m *MockDataManager) CreateUserInfo(userId string, userName string) (userInfo dbModel.UserInfo, err error) {
	userInfo = dbModel.UserInfo{
		Id:       userId,
		UserName: userName,
	}
	if err = database.GetSqlDb().Orm.Create(&userInfo).Error; err != nil {
		return
	}

	return
}
