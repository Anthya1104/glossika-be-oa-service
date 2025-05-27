package data

import (
	"github.com/Anthya1104/gin-base-service/internal/app/database"
	dbModel "github.com/Anthya1104/gin-base-service/internal/app/model/db"
)

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
