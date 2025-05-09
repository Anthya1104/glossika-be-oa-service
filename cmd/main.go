package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Anthya1104/gin-base-service/internal/app/router"
	"github.com/spf13/viper"
)

func main() {

	// init viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	port := viper.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	host := viper.GetString("server.host")
	if host == "" {
		host = "localhost"
	}

	r := router.SetupRouter()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s", err)
		}
	}()
	log.Printf("Server is running at http://%s:%d", host, port)

	// Wait for a signal to shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// 5 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
