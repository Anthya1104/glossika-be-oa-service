package database

import (
	"context"

	dbModel "github.com/Anthya1104/glossika-be-oa-service/internal/app/model/db"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/errcode"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/orm"
	"gorm.io/gorm"
)

type Repo interface {
	// get users
	GetUserInfoByUserId(ctx context.Context, userId string) (userInfo dbModel.UserInfo, wrapErr errcode.WrapErr)
}

var sqlDbInstance *SqlDb

func GetSqlDb() *SqlDb {
	return sqlDbInstance
}

type SqlDb struct {
	Orm  *gorm.DB
	Repo Repo
}

func (s *SqlDb) CloseConnection() {
	if sqlDB, err := s.Orm.DB(); err != nil {
		log.L.Errorf("orm db connection close fail: %v", err)
	} else {
		sqlDB.Close()
	}
}

func NewSqlDb(conf orm.Config) error {
	db, err := orm.SetupMysqlDb(conf)
	if err != nil {
		return err
	}

	sqlDbInstance = &SqlDb{
		Orm:  db,
		Repo: &SqlDb{Orm: db},
	}

	return nil
}

func AutoMigrate(db *gorm.DB) error {
	models := []interface{}{
		&dbModel.UserInfo{},
	}
	return orm.AutoMigrate(db, models)

}

// Only use for integration test
func DropTables(db *gorm.DB) error {
	models := []interface{}{
		&dbModel.UserInfo{},
	}
	return orm.DropTables(db, models)
}
