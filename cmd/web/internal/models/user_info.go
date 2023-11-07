// Filename: internal/models/user_info.go
package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

type User struct {
	UserID      int64
	Fname       string
	Lname       string
	Email       string
	Address     string
	PhoneNumber string
	Password    string
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(fname, lname, email, address, phone_number, password string) error {
	//let's Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `
			INSERT INTO users_info(fname, lname, email, addres, phone_number, passwrd)
			VALUES($1, $2, $3, $4, $5, $6)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, fname, lname, email, address, phone_number, hashedPassword)
	if err != nil {
		switch {
		case err.Error() == `pgx: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

// This was added
func (m *UserModel) Authenticate(email, password string) (int, error) {
	log.Println(password)
	var id int
	var hashedPassword []byte
	//create if there is a row in the table for the email provided
	query := `
			SELECT id, passwrd
			FROM users_info
			WHERE email = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	} //handling error
	//the user does exist
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	log.Println("Leave authenticate")
	//password is correct
	return id, nil
}
