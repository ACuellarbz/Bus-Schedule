// main file

package main

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ACuellarbz/3162/internal/models"
	"github.com/alexedwards/scs/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type application struct {
	errorLog        *log.Logger
	infoLog         *log.Logger
	user_info       models.UserModel
	bus_company     models.BusCompanyModel
	bus_schedule    models.BusScheduleModel
	locations       models.LocationModel
	route           models.RouteModel
	seat            models.SeatModel
	ticket_orders   models.TicketOrdersModel
	sessionsManager *scs.SessionManager
}

func main() {
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("COMPUTE_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}
	//create instances of errorLog & infoLog
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// setup a new session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = false                  //false if the cookies aren't secure
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode //Same site

	//share data across handlers
	app := &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		user_info:       models.UserModel{DB: db},
		bus_company:     models.BusCompanyModel{DB: db},
		bus_schedule:    models.BusScheduleModel{DB: db},
		locations:       models.LocationModel{DB: db},
		route:           models.RouteModel{DB: db},
		seat:            models.SeatModel{DB: db},
		ticket_orders:   models.TicketOrdersModel{DB: db},
		sessionsManager: sessionManager,
	}

	defer db.Close()
	log.Println("Database connection pool established")

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	log.Printf("Starting server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Get a database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	//use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //specify timeout
	defer cancel()                                                          //+Defer attached to a function and executes as the last thing
	//lets ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
