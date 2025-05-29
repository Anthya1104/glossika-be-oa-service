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
