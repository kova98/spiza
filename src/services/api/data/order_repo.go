package data

import (
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderRepo(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db}
}

type OrderItem struct {
	Id      int64   `json:"id,omitempty"`
	Name    string  `json:"name"`
	Order   int32   `json:"order" db:"order_num"`
	Price   float64 `json:"price"`
	Image   string  `json:"image"`
	OrderId int64   `db:"order_id"`
}

type OrderWithItems struct {
	Id          int64       `json:"id"`
	UserId      int64       `json:"user_id" db:"user_id"`
	Status      int64       `json:"status"`
	DateCreated time.Time   `json:"date_created" db:"date_created"`
	Items       []OrderItem `json:"items" `
}

func (r *OrderRepo) CreateOrder(o *Order) (*Order, error) {
	var created Order
	sql := `INSERT INTO orders (restaurant_id, user_id, destination_id, items) VALUES($1, $2, $3, $4) 
		    RETURNING id, user_id, restaurant_id, destination_id, status, date_created, items;`
	if err := r.db.Get(&created, sql, o.RestaurantId, o.UserId, o.DestinationId, SqlArrayValue(o.Items)); err != nil {
		return nil, err
	}
	return &created, nil
}

func (r *OrderRepo) GetOrder(id int64) (*Order, error) {
	var order Order
	sql := `SELECT * FROM orders WHERE id = $1;`
	if err := r.db.Get(&order, sql, id); err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepo) GetOrders(restaurantId int64) ([]OrderWithItems, error) {
	var orders []OrderWithItems
	oQuery := "SELECT id, user_id, status, date_created FROM orders WHERE restaurant_id = $1 ORDER BY date_created DESC;"
	if err := r.db.Select(&orders, oQuery, restaurantId); err != nil {
		return nil, err
	}

	var items []OrderItem
	sql := `SELECT o.id as order_id, i.id, i.name, i.order_num, i.price
			FROM orders o
			JOIN items i ON i.id = ANY(o.items)
			WHERE o.restaurant_id = $1`
	if err := r.db.Select(&items, sql, restaurantId); err != nil {
		return nil, err
	}

	for i, order := range orders {
		orders[i].Items = []OrderItem{}
		for _, item := range items {
			if order.Id == item.OrderId {
				orders[i].Items = append(orders[i].Items, item)
			}
		}
	}

	return orders, nil
}

func (r *OrderRepo) UpdateOrderStatus(orderId int64, status int) error {
	sql := `UPDATE orders SET status = $1 WHERE id = $2;`
	_, err := r.db.Exec(sql, status, orderId)
	return err
}

func (r *OrderRepo) SetCourier(orderId int64, courierId int64) error {
	sql := `UPDATE orders SET courier_id = $1 WHERE id = $2;`
	_, err := r.db.Exec(sql, courierId, orderId)
	return err
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
