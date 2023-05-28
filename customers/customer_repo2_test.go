package customers

import "testing"

import (
	"context"

	"github.com/sivaprasadreddy/learn-go-testing/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepositoryUsingTestUtils(t *testing.T) {
	ctx := context.Background()

	container, closeFn, err := testsupport.GetPostgresContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(closeFn)

	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)

	conn := testsupport.GetDb(connStr)
	customerRepo := NewCustomerRepo(conn)

	customers, err := customerRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}
