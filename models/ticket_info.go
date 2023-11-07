// Filename: internal/models/seats
package models

import (
	"context"
	"database/sql"
	"time"
)

type Seats struct {
	TicketInfoID    int64
	ScheduleID      int64
	TicketPrice     int64 //Need to Ask Question on How to Reference
	NumberofTickets int64
}

type SeatModel struct {
	DB *sql.DB
}

// Code to access the database
func (m *SeatModel) Get() (*Seats, error) {
	var t Seats

	statement := `
				SELECT id, seat_name, type_of_seat
				FROM ticket_info
				LIMIT 1
				`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&t.TicketInfoID, &t.ScheduleID, &t.TicketPrice, &t.NumberofTickets)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
