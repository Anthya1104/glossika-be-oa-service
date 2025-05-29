package db

import (
	"time"

	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
)

func (UserRecommendation) TableName() string {
	return string(config.TableNameUserRecommendation)
}

type UserRecommendation struct {
	ID        uint    `gorm:"primaryKey;autoIncrement;column:id"`
	UserID    uint    `gorm:"column:user_id"`
	ProductID uint    `gorm:"column:product_id"`
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time
}
