package models

import (
	"context"
	"encoding/json"
	"nats_service/internal/database/postgres"

	"github.com/jackc/pgx/v4"
)

type Order struct {
	OrderUid          string        `json:"order_uid"`
	TrackNumber       string        `json:"track_number"`
	Entry             string        `json:"entry"`
	Delivery          OrderDelivery `json:"delivery"`
	Payment           OrderPayment  `json:"payment"`
	Items             []OrderItem   `json:"items"`
	Locale            string        `json:"locale"`
	InternalSignature string        `json:"internal_signature"`
	CustomerId        string        `json:"customer_id"`
	DeliveryService   string        `json:"delivery_service"`
	Shardkey          string        `json:"shardkey"`
	SmId              int64         `json:"sm_id"`
	DateCreated       string        `json:"date_created"`
	OOFShard          string        `json:"oof_shard"`
}

func GetAllOrders(db *postgres.Database) ([]*Order, error) {
	var orders []*Order

	rows, err := db.Query(context.Background(), "SELECT get_all_orders()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		order, err := orderFromRow(rows)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, rows.Err()
}

func orderFromRow(row pgx.Row) (*Order, error) {
	var order Order
	var jsonData []byte

	if err := row.Scan(&jsonData); err == pgx.ErrNoRows {
		return nil, nil
	}
	if err := json.Unmarshal(jsonData, &order); err != nil {
		return nil, err
	}

	return &order, nil
}
