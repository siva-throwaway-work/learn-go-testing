package orders

import (
	"context"
	"testing"

	"github.com/sivaprasadreddy/learn-go-testing/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestOrderPublish(t *testing.T) {
	ctx := context.Background()

	conn := testsupport.GetDb(dbConnString)
	orderRepo := NewOrderRepo(conn)
	orderSync := OrderPublisher{repo: &orderRepo}

	err := orderSync.publishToRedis(ctx)
	assert.NoError(t, err)
}
