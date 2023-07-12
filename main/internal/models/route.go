// Filename: internal/models/routes
package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Route struct {
	RouteID       int64
	BeginningID   int64
	DestinationID int64
	TypeTrip      string
	DepartTime    string
	ArrivalTIme   string
}

type RouteModel struct {
	DB *sql.DB
}

// Code to access the database

func (m *RouteModel) Insert(route_id string, beginning string, destination string, type_trip string, depart_time string, arrival_time string) (int64, error) {
	var id int64

	statement :=
		`
			INSERT INTO routes(id, beginning_location_id, destination_location_id, type_of_trip, bus_departutre_time, bus_arrival_time)
			VALUES($1, $2, $3, $4, $5, $6)
			RETURNING ID
		`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, route_id, beginning, destination, type_trip, depart_time, arrival_time).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *RouteModel) Delete(route_id string) error {
	//Statement that will delete record
	statement := `
	DELETE FROM routes
	WHERE id = $1
	`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := m.DB.ExecContext(ctx, statement, route_id)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Code to access the database
func (m *RouteModel) Get() (*Route, error) {
	var b Route

	statement := `
				SELECT id, beginning_location_id, destination_location_id, type_of_trip, bus_departure_time, bus_arrival_time
				FROM routes
				`
	//timeout for DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&b.RouteID, &b.BeginningID, &b.DestinationID, &b.TypeTrip, &b.DepartTime, &b.ArrivalTIme)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

// Credit: Interpretation of Hipolito's Read Function
func (m *RouteModel) SearchRecord(route_id string) ([]*Route, string, error) {
	//SQL statement
	statement :=
		`
			SELECT id, beginning_location_id, destination_location_id, type_of_trip, bus_departure_time, bus_arrival_time
			FROM routes
			WHERE id = $1
		`
	failure := "program failed"
	data, err := m.DB.Query(statement, route_id)
	if err != nil {
		return nil, failure, err
	}
	defer data.Close()
	log.Println(data)

	routes := []*Route{}

	data.Next()

	route := &Route{}
	err = data.Scan(&route.RouteID, &route.BeginningID, &route.DestinationID, &route.TypeTrip, &route.DepartTime, &route.ArrivalTIme)

	routes = append(routes, route)
	if err != nil {
		fmt.Println(err)
		return nil, failure, err
	}
	if err = data.Err(); err != nil {
		return nil, failure, err
	}
	return routes, route_id, nil
}

func (m *RouteModel) Update(route_id string, beginning string, destination string, type_trip string, depart_time string, arrival_time string) error {
	statement := `
			UPDATE routes
			SET id = $1, beginning_location_id = $2, destination_location_id = $3, type_of_trip = $4, bus_departure_time = $5, bus_arrival_time = $6
			WHERE id = $7
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//sets the timeout for the DB connection, passes the statements and associates the arguements with the place holders in the SQL ($1, $2 etc)
	_, err := m.DB.ExecContext(ctx, statement, route_id, beginning, destination, type_trip, depart_time, arrival_time)

	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}
