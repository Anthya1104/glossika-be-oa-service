package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Anthya1104/gin-base-service/pkg/log"
	"github.com/spf13/viper"
)

type LogLevel string

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelWarning LogLevel = "warn"
	LogLevelError   LogLevel = "error"
)

type DevelopEnv string

const (
	deployEnvDevelop    DevelopEnv = "develop"
	deployEnvStage      DevelopEnv = "stage"
	deployEnvProduction DevelopEnv = "production"
)

var EnvVariable envVariable

type envVariable struct {
	Version           string
	Host              string
	Port              string
	DeployEnvironment string
	LogLevel          string

	SQLHost     string
	SQLPort     string
	SQLDatabase string
	SQLUsername string
	SQLPassword string
}

func Validate() (err error) {
	port, err := strconv.ParseUint(EnvVariable.Port, 10, 16)
	if err != nil || port <= 0 || port > uint64(65535) {
		err = errors.New("required environment variable \"GO_HTTP_PORT\" should be 0~65535")
		return
	}
	if l := strings.ToLower(EnvVariable.LogLevel); LogLevel(l) != LogLevelError && LogLevel(l) != LogLevelDebug && LogLevel(l) != LogLevelWarning && LogLevel(l) != LogLevelInfo {
		err = errors.New("required environment variable \"LOG_LEVEL\" should be \"ERROR|DEBUG|WARN|INFO\"")
		return
	}
	if d := strings.ToLower(EnvVariable.DeployEnvironment); DevelopEnv(d) != deployEnvDevelop && DevelopEnv(d) != deployEnvStage && DevelopEnv(d) != deployEnvProduction {
		err = errors.New("required environment variable \"DEPLOY_ENVIRONMENT\" should be \"DEVELOP|STAGE|PRODUCTION\"")
		return
	}

	return
}

func IsProduction() bool {
	return strings.ToLower(EnvVariable.DeployEnvironment) == string(deployEnvProduction)
}

func Setup() error {
	configName := os.Getenv("CONFIG_NAME")
	if configName == "" {
		configName = "config"
	}
	log.L.Infof("using config file: %s", configName)

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("config error: %w", err)
		}
	}

	viper.AutomaticEnv()

	log.L.Infof("viper settings: %+v\n", viper.AllSettings())

	requiredEnvs := []string{
		"server.port",
		"server.host",
		"database.host",
		"database.port",
		"database.name",
		"database.user",
		"database.password",
	}

	// for test env init
	viper.BindEnv("database.host", "SQL_HOST")
	viper.BindEnv("database.port", "SQL_PORT")
	viper.BindEnv("database.name", "SQL_DATABASE")
	viper.BindEnv("database.user", "SQL_USERNAME")
	viper.BindEnv("database.password", "SQL_PASSWORD")

	for _, env := range requiredEnvs {
		if viper.GetString(env) == "" {
			return fmt.Errorf("required environment variable %s is not set", env)
		}
	}

	// init env variable
	EnvVariable = envVariable{
		Version:           viper.GetString("version"),
		Host:              viper.GetString("server.host"),
		Port:              viper.GetString("server.port"),
		DeployEnvironment: viper.GetString("server.env"),
		LogLevel:          viper.GetString("logging.level"),
		SQLHost:           viper.GetString("database.host"),
		SQLPort:           viper.GetString("database.port"),
		SQLDatabase:       viper.GetString("database.name"),
		SQLUsername:       viper.GetString("database.user"),
		SQLPassword:       viper.GetString("database.password"),
	}

	// set default value
	if EnvVariable.Version == "" {
		EnvVariable.Version = "1.0.0"
	}
	if EnvVariable.DeployEnvironment == "" {
		EnvVariable.DeployEnvironment = "DEVELOP"
	}
	if EnvVariable.LogLevel == "" {
		EnvVariable.LogLevel = "INFO"
	}

	if err := Validate(); err != nil {
		fmt.Printf("config validate fail, %+v\n", err)
		return err
	}

	return nil
}
