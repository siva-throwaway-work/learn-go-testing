package customers

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Customer struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomerRepo struct {
	conn *pgx.Conn
}

func NewCustomerRepo(conn *pgx.Conn) CustomerRepo {
	return CustomerRepo{
		conn: conn,
	}
}

func (p CustomerRepo) GetAll(ctx context.Context) ([]Customer, error) {
	rows, err := p.conn.Query(ctx, `SELECT id, name, email FROM customers`)
	if err != nil {
		return nil, err
	}
	var customers []Customer

	defer rows.Close()
	for rows.Next() {
		var customer = Customer{}
		err = rows.Scan(&customer.Id, &customer.Name, &customer.Email)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (p CustomerRepo) GetByEmail(ctx context.Context, email string) (Customer, error) {
	var customer Customer
	err := p.conn.QueryRow(ctx, `SELECT id, name,email FROM customers WHERE email = $1`, email).Scan(
		&customer.Id, &customer.Name, &customer.Email)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}
