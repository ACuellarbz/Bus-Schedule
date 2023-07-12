// Filename: cmd/web/data.go
package main

import "github.com/ACuellarbz/3162/internal/models"

type templateData struct {
	Schedule     *models.Route
	Flash        string
	ScheduleByte []*models.Route //used to hold byte data I guess
	CSRFTOKEN    string          // Added for authentication
}
