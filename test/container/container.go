package container

import (
	"context"
	"sync"

	"github.com/Anthya1104/glossika-be-oa-service/pkg/log"
)

type TestContainers struct {
	sqlContainer *MySQLContainer
}

func (t *TestContainers) TearDownContainers(ctx context.Context) {
	t.sqlContainer.Terminate(ctx)
}

func NewTestContainers(ctx context.Context) TestContainers {
	testContainers := TestContainers{}

	wg := new(sync.WaitGroup)

	wg.Add(1)
	var setupSQLErr error
	go func() {
		defer wg.Done()
		testContainers.sqlContainer, setupSQLErr = SetupMySQLTestContainer(ctx)
		if setupSQLErr != nil {
			log.L.Fatalf("setupSQL Fail: %v", setupSQLErr)
		}
	}()

	wg.Wait()
	return testContainers
}
