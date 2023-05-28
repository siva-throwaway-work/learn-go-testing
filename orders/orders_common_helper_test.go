package orders

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/sivaprasadreddy/learn-go-testing/testsupport"
)

//A common test file with TestMain(..) which executes once per entire package
// TestMain(..) will be called only once per package,
// All the containers will be started and be used by any test within this package.

// Issue:
// Having TestMain(..) behaviour per-package instead of per-test file results in not having more control.

var dbConnString string
var redisConnString string

func TestMain(m *testing.M) {
	ctx := context.Background()
	pgContainer, dbCloseFn, err := testsupport.GetPostgresContainer(ctx)
	if err != nil {
		log.Fatalf("failed to setup Postgres container. Error: %v", err)
	}
	dbConnString, err = pgContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get Postgres connectionstring. Error: %v", err)
	}
	defer dbCloseFn()

	redisContainer, redisCloseFn, err := testsupport.GetRedisContainer(ctx)
	if err != nil {
		log.Fatalf("failed to setup Redis container. Error: %v", err)
	}
	redisConnString, err = redisContainer.ConnectionString(ctx)
	if err != nil {
		log.Fatalf("failed to get Redis connectionstring. Error: %v", err)
	}
	log.Printf("Redis ConnectionString: %s", redisConnString)
	defer redisCloseFn()

	os.Exit(m.Run())
}
