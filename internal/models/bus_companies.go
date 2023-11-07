// Filename: internal/models/bus_schedule.go
package models

import (
	"context"
	"database/sql"
	"time"
)

type BusCompany struct {
	CompanyID   int64
	CompanyName string
}

type BusCompanyModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *BusCompanyModel) Get() (*BusCompany, error) {
	var c BusCompany

	statement := `
				SELECT id, company_name
				FROM bus_companies
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&c.CompanyID, &c.CompanyName)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
