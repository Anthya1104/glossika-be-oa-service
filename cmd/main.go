package main

import (
	"fmt"
	"log"

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

	r := router.SetupRouter()

	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
