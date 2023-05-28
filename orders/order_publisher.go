package orders

import (
	"context"
	"fmt"
)

type OrderPublisher struct {
	repo *OrderRepo
}

func (c OrderPublisher) publishToRedis(ctx context.Context) error {
	orders, err := c.repo.GetAll(ctx)
	if err != nil {
		return err
	}
	for _, order := range orders {
		fmt.Printf("Send Order : %d info to Redis \n", order.Id)
	}
	return nil
}
