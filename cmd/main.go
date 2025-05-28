package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anthya1104/glossika-be-oa-service/internal/app/database"
	"github.com/Anthya1104/glossika-be-oa-service/internal/app/router"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/config"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/orm"
	"github.com/Anthya1104/glossika-be-oa-service/pkg/redis"
)

func main() {

	// init env
	if err := config.Setup(); err != nil {
		log.L.Fatal(err)
	}

	// TODO: only skip pswd in demo, should not be empty in productive environment
	// setup redis
	addrString := fmt.Sprintf("%s:%s", config.EnvVariable.RedisHost, config.EnvVariable.RedisPort)
	redis.InitRedis(addrString, "", 0)
	if err := redis.Ping(); err != nil {
		log.L.Fatalf("Redis connect failed: %v", err)
	}
	// setup db
	if err := database.NewSqlDb(orm.Config{
		Host:     config.EnvVariable.SQLHost,
		Port:     config.EnvVariable.SQLPort,
		Database: config.EnvVariable.SQLDatabase,
		Username: config.EnvVariable.SQLUsername,
		Password: config.EnvVariable.SQLPassword,
	}); err != nil {
		log.L.Fatalf("failed to connect database: %v", err)
	}

	if err := database.AutoMigrate(database.GetSqlDb().Orm); err != nil {
		log.L.Fatalf("auto migrate failed: %v", err)
	}

	r := router.SetupRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.EnvVariable.Port),
		Handler:      r,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.L.Fatalf("listen: %s", err)
		}
	}()

	// Wait for a signal to shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.L.Debug("Shutdown Server ...")

	// 5 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.L.Fatal("Server Shutdown:", err)
	}

	log.L.Debug("Server exiting")
}
