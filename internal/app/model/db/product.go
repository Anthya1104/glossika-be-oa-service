package db

import "github.com/Anthya1104/glossika-be-oa-service/pkg/config"

func (Product) TableName() string {
	return string(config.TableNameProduct)
}

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement;column:id"`
	ProductID   string  `gorm:"type:varchar(20);uniqueIndex;column:product_id"`
	Name        string  `gorm:"type:varchar(20);column:product_name"`
	Description string  `gorm:"type:varchar(256);column:description"`
	Price       float64 `gorm:"type:decimal(10,2);column:price"`
}
