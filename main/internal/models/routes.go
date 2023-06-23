// Filename: internal/models/routes
package models

import (
	"context"
	"database/sql"
	"time"
)

type Route struct {
	MileageID      int64
	BeginningID       int64
	DestinationID       int64
	TotalMiles       int64
}

type RouteModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *RouteModel) Get() (*Route, error) {
	var o Route

	
	statement := `
				SELECT id, route_name, number_of_miles, total_cost, number_of_tickets_available
				FROM route
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&o.MileageID, &o.BeginningID, &o.DestinationID, &o.TotalMiles)
	if err != nil {
		return nil, err
	}
	return &o, nil
}
