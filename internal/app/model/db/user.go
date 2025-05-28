package db

import (
	"time"

	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
)

func (User) TableName() string {
	return string(config.TableNameUser)
}

type User struct {
	ID          uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Email       string `gorm:"type:varchar(254);uniqueIndex;not null;column:user_email"`
	Password    string `gorm:"type:varchar(60);not null;column:user_password"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsActivated bool `gorm:"type:tinyint(1);not null;default:false;column:is_activated"`
}

// TODO: remove redundant sample after feature completed
func (UserInfo) TableName() string {
	return string(config.TableNameUserInfo)
}

type UserInfo struct {
	Id       string `gorm:"type:varchar(50);primary_key;column:id"`
	UserName string `gorm:"type:varchar(50);not null;column:user_name"`
}
