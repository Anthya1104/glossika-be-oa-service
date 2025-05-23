package orm

import (
	"fmt"
	"strings"

	"github.com/Anthya1104/gin-base-service/pkg/config"
	"github.com/Anthya1104/gin-base-service/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

// TODO: 看point-service 的 gorm.go 整理下面, log 應該要稍微改一下，或是先抽掉
func SetupMysqlDb(cfg Config) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
	// 	cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	// return gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	enableDebug := strings.EqualFold(config.EnvVariable.LogLevel, string(config.LogLevelDebug))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: log.NewGormLogger(enableDebug),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}
