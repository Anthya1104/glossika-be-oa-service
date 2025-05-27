package container

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MySQLContainer struct {
	testcontainers.Container
	URI string
}

func SetupMySQLTestContainer(ctx context.Context) (*MySQLContainer, error) {
	const (
		image          = "mysql:8.0"
		driver         = "mysql"
		username       = "testuser"
		rootpaswd      = "root123"
		password       = "testpwd"
		dbName         = "testdb"
		port           = "3306/tcp"
		startupTimeout = time.Second * 60
		imagePlatform  = "linux/amd64"
	)

	req := testcontainers.ContainerRequest{
		Image: image,
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": rootpaswd,
			"MYSQL_DATABASE":      dbName,
			"MYSQL_USER":          username,
			"MYSQL_PASSWORD":      password,
		},
		ExposedPorts: []string{port},
		WaitingFor: wait.ForSQL(port, driver, func(host string, port nat.Port) string {
			return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port.Port(), dbName)
		}).WithStartupTimeout(startupTimeout),
		AutoRemove:    true,
		ImagePlatform: imagePlatform,
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		return nil, err
	}

	hostIP, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	setEnvVars(hostIP, mappedPort.Port(), dbName, username, password)

	uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", hostIP, username, password, dbName, mappedPort.Port())

	return &MySQLContainer{Container: container, URI: uri}, nil
}

func setEnvVars(host, port, dbName, username, password string) {
	os.Setenv("SQL_HOST", host)
	os.Setenv("SQL_PORT", port)
	os.Setenv("SQL_DATABASE", dbName)
	os.Setenv("SQL_USERNAME", username)
	os.Setenv("SQL_PASSWORD", password)
}
