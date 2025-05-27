package db

import (
	"github.com/Anthya1104/gin-base-service/pkg/config"
)

func (UserInfo) TableName() string {
	return string(config.TableNameUserInfo)
}

type UserInfo struct {
	Id       string `gorm:"type:varchar(50);primary_key;column:id"`
	UserName string `gorm:"type:varchar(50);not null;column:user_name"`
}
