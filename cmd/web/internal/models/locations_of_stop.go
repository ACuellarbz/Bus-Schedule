// Filename: internal/models/locations_of_stop.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type Location struct {
	LocationID      int64
	LocationName 	string
}

type LocationModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *LocationModel) Get() (*Location, error) {
	var l Location

	statement := `
				SELECT id, location
				FROM locations_of_stop
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&l.LocationID, &l.LocationName)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
