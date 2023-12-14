package data

import (
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db}
}

func (r *OrderRepo) CreateOrder(o *Order) (int64, error) {
	var id int64
	sql := "INSERT INTO orders (restaurant_id, user_id, items) VALUES($1, $2, $3) RETURNING id;"
	if err := r.db.Get(&id, sql, o.RestaurantId, o.UserId, SqlArrayValue(o.Items)); err != nil {
		return 0, err
	}
	return id, nil
}

func SqlArrayValue(ids []int64) string {
	b := strings.Builder{}
	b.WriteString("{")
	for i, item := range ids {
		b.WriteString(strconv.FormatInt(item, 10))
		if i < len(ids)-1 {
			b.WriteString(",")
		}
	}
	b.WriteString("}")
	return b.String()
}
