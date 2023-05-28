package orders

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Order struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}

type OrderRepo struct {
	conn *pgx.Conn
}

func NewOrderRepo(conn *pgx.Conn) OrderRepo {
	return OrderRepo{
		conn: conn,
	}
}

func (p OrderRepo) GetAll(ctx context.Context) ([]Order, error) {
	rows, err := p.conn.Query(ctx, `SELECT id, amount FROM orders`)
	if err != nil {
		return nil, err
	}
	var orders []Order

	defer rows.Close()
	for rows.Next() {
		var order = Order{}
		err = rows.Scan(&order.Id, &order.Amount)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
