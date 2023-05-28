package customers

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Preference struct {
	Id        int  `json:"id"`
	CustId    int  `json:"cust_id"`
	Subscribe bool `json:"subscribe"`
}

type PreferenceRepo struct {
	conn *pgx.Conn
}

func NewPreferenceRepo(conn *pgx.Conn) PreferenceRepo {
	return PreferenceRepo{
		conn: conn,
	}
}

func (p PreferenceRepo) CreatePreference(ctx context.Context, preference Preference) (Preference, error) {
	err := p.conn.QueryRow(ctx,
		"INSERT INTO preferences (id, cust_id, subscribed) VALUES ($1, $2, $3) RETURNING id",
		preference.Id, preference.CustId, preference.Subscribe).Scan(&preference.Id)
	return preference, err
}

func (p PreferenceRepo) GetByCustId(ctx context.Context, custId int) (Preference, error) {
	var preference Preference
	err := p.conn.QueryRow(ctx, `SELECT id, cust_id,subscribed FROM preferences WHERE cust_id = $1`, custId).Scan(
		&preference.Id, &preference.CustId, &preference.Subscribe)
	if err != nil {
		return Preference{}, err
	}
	return preference, nil
}
