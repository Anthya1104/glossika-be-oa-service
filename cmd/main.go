package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anthya1104/gin-base-service/internal/app/router"
	"github.com/Anthya1104/gin-base-service/pkg/config"
	"github.com/Anthya1104/gin-base-service/pkg/log"
)

func main() {

	// init env
	if err := config.Setup(); err != nil {
		log.L.Fatal(err)
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
