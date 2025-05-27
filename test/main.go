package test

import (
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/router"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/orm"
)

func Setup() {
	if err := config.Setup(); err != nil {
		log.L.Fatal(err)
	}

	if err := database.NewSqlDb(orm.Config{
		Host:     config.EnvVariable.SQLHost,
		Port:     config.EnvVariable.SQLPort,
		Database: config.EnvVariable.SQLDatabase,
		Username: config.EnvVariable.SQLUsername,
		Password: config.EnvVariable.SQLPassword,
	}); err != nil {
		log.L.Fatal(err)
	}

	if err := log.Setup(config.EnvVariable.LogLevel); err != nil {
		log.L.Fatal(err)
	}

	if err := router.Setup(); err != nil {
		log.L.Fatal(err)
	}
}
