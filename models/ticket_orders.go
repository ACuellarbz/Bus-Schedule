// Filename: internal/models/ticket_orders.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type TicketOrders struct {
	TicketOrdersID int64
	TicketInfoID   int64
	UserID         int64
	PurchasedAt    time.Time
}

type TicketOrdersModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *TicketOrdersModel) Get() (*TicketOrders, error) {
	var t TicketOrders

	statement := `
				SELECT id, ticket_id, user_id
				FROM ticket_orders
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&t.TicketOrdersID, &t.TicketInfoID, &t.UserID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
