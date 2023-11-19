// Filename: cmd/web/data.go
package main

import "github.com/ACuellarbz/Bus-Schedule/internal/models"

type templateData struct {
	Schedule     *models.Route
	Flash        string
	ScheduleByte []*models.Route //used to hold byte data I guess
	CSRFTOKEN    string          // Added for authentication
}
