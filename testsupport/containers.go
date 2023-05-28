package testsupport

import (
	"context"
	"log"
	"path/filepath"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/modules/redis"
	"github.com/testcontainers/testcontainers-go/wait"
)

/**
Can't share test helper code from *_test.go files.
https://stackoverflow.com/questions/56404355/how-do-i-package-golang-test-helper-code

The alternative approach is to put them in non _test.go files under a package that indicates its purpose
is for testing. Ex: httptest, zaptest.

Hope no one will call these public functions from production code accidentally, and pray :-)

Assumption: If these helpers are being called only from test code
then these test helpers won't be packaged in final production binary.
We Need to verify that.
*/

func GetPostgresContainer(ctx context.Context) (*postgres.PostgresContainer, func(), error) {
	container, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.2-alpine"),
		postgres.WithInitScripts(filepath.Join("../testdata", "initdb.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanFn := func() {
		if err := container.Terminate(ctx); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}

	return container, cleanFn, nil
}

func GetRedisContainer(ctx context.Context) (*redis.RedisContainer, func(), error) {
	container, err := redis.RunContainer(ctx,
		testcontainers.WithImage("redis:latest"),
	)
	if err != nil {
		return nil, nil, err
	}

	cleanFn := func() {
		if err := container.Terminate(ctx); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}

	return container, cleanFn, nil
}
