package customers

import (
	"context"
	"testing"

	"github.com/sivaprasadreddy/learn-go-testing/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	ctx := context.Background()

	conn := testsupport.GetDb(dbConnString)
	customerRepo := NewCustomerRepo(conn)

	customers, err := customerRepo.GetAll(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, customers)
}
