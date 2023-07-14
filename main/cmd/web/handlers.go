package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/ACuellarbz/3162/internal/models"
	"github.com/justinas/nosurf"
)

var dataStore = struct {
	sync.RWMutex
	data map[string]int64
}{data: make(map[string]int64)}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.tmpl", nil)

}

func (app *application) tickets(w http.ResponseWriter, r *http.Request) {

}

// Read Implementation
func (app *application) scheduleShow(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered Schedule")
	schedule, err := app.route.Get()
	if err != nil {
		log.Println(err)
		return
	}
	data := &templateData{
		Schedule:  schedule,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "schedule.page.tmpl", data)
}
func (app *application) scheduleFormShow(w http.ResponseWriter, r *http.Request) {
	data := &templateData{
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "schedule.add.tmpl", data)

}

// POST METHOD implementation of Create
// NEEDS TO BE FIXED
func (app *application) scheduleFormSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	begin_location := r.PostForm.Get("begin_id")
	destin_location := r.PostForm.Get("destination_id")
	trip_type := r.PostForm.Get("type_trip")
	departure := r.PostForm.Get("departure_time")
	arrivals := r.PostForm.Get("arrival")

	//Make Adjustments
	if trip_type != "Regular" && trip_type != "Express" {
		log.Println("Not Correct Type")
	}

	//SWITCH FOR begin_location
	switch {
	case begin_location == "Belize":
		begin_location = "01"
	case begin_location == "Belmopan":
		begin_location = "02"
	case begin_location == "Cayo":
		begin_location = "03"
	case begin_location == "Punta Gorda":
		begin_location = "04"
	case begin_location == "Dangriga":
		begin_location = "05"
	case begin_location == "Stann Creek":
		begin_location = "06"
	case begin_location == "Corozal":
		begin_location = "07"
	case begin_location == "Orange Walk":
		begin_location = "08"
	default:
		log.Println("SWITCH FAILED")
	}

	//SWITCH FOR destin_location
	switch {
	case destin_location == "Belize":
		destin_location = "01"
	case destin_location == "Belmopan":
		destin_location = "02"
	case destin_location == "Cayo":
		destin_location = "03"
	case destin_location == "Punta Gorda":
		destin_location = "04"
	case destin_location == "Dangriga":
		destin_location = "05"
	case destin_location == "Stann Creek":
		destin_location = "06"
	case destin_location == "Corozal":
		destin_location = "07"
	case destin_location == "Orange Walk":
		destin_location = "08"
	default:
		log.Println("SWITCH FAILED")
	}

	log.Printf("%s %s %s %s %s\n", begin_location, destin_location, trip_type, departure, arrivals)

	_, err = app.route.Insert(id, begin_location, destin_location, trip_type, departure, arrivals)
	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
func (app *application) login(w http.ResponseWriter, r *http.Request) {
	//remove the entry from the session manager
	flash := app.sessionsManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "login.page.tmpl", data)
}

func (app *application) loginSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	//write the data to the table
	id, err := app.user_info.Authenticate(email, password)
	log.Println(err)
	if err != nil {

		if errors.Is(err, models.ErrInvalidCredentials) {
			RenderTemplate(w, "login.page.tmpl", nil)
		}
		return
	}
	//add the user to the session cookie
	err = app.sessionsManager.RenewToken(r.Context())
	if err != nil {
		return
	}
	// add an authenticate entry
	app.sessionsManager.Put(r.Context(), "authenticatedUserID", id)
	http.Redirect(w, r, "/schedule", http.StatusSeeOther)
}

func (app *application) signup(w http.ResponseWriter, r *http.Request) {
	flash := app.sessionsManager.PopString(r.Context(), "flash")
	data := &templateData{
		Flash:     flash,
		CSRFTOKEN: nosurf.Token(r), //added for authentication
	}
	RenderTemplate(w, "signup.page.tmpl", data)
}

func (app *application) signupSubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fname := r.PostForm.Get("fname")
	lname := r.PostForm.Get("lname")
	email := r.PostForm.Get("email")
	address := r.PostForm.Get("addres")
	phone_number := r.PostForm.Get("phone_number")
	password := r.PostForm.Get("passwrd")

	//write the data to the table
	err := app.user_info.Insert(fname, lname, email, address, phone_number, password)
	log.Println(err)
	if err != nil {

		if errors.Is(err, models.ErrDuplicateEmail) {
			app.sessionsManager.Put(r.Context(), "flash", "Email Already Registered")
			RenderTemplate(w, "signup.page.tmpl", nil)
		}
	}
	app.sessionsManager.Put(r.Context(), "flash", "Signup was successfil")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) logoutSubmit(w http.ResponseWriter, r *http.Request) {
	//remove the entry from the session manager
	err := app.sessionsManager.RenewToken(r.Context())
	if err != nil {
		return
	}

	app.sessionsManager.Remove(r.Context(), "authenticatedUserID")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Displays Update Request Page
func (app *application) updateScheduleShow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "schedule.update.request.tmpl", nil)
}

// POST METHOD Implementation for Update
func (app *application) updateSchedule(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	info, schedule_id, err := app.route.SearchRecord(id)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	schedule, err := strconv.Atoi(schedule_id)
	if err != nil {
		// ... handle error
		panic(err)
	}

	data := &templateData{
		ScheduleByte: info,
	}

	ts, err := template.ParseFiles("./ui/html/schedule.update.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	dataStore.Lock()
	dataStore.data["key"] = int64(schedule)
	dataStore.Unlock()

	log.Println(data)
	err = ts.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}
func (app *application) updateRecords(w http.ResponseWriter, r *http.Request) {
	log.Println("Im inside updateRecords")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	begin_location := r.PostForm.Get("begin_id")
	destin_location := r.PostForm.Get("destination_id")
	trip_type := r.PostForm.Get("type_trip")
	departure := r.PostForm.Get("departure_time")
	arrivals := r.PostForm.Get("arrival")
	err = app.route.Update(id, begin_location, destin_location, trip_type, departure, arrivals)
	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (app *application) deleteRouteShow(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "schedule.delete.tmpl", nil)

}
func (app *application) deleteRoute(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
	id := r.PostForm.Get("id")
	err = app.route.Delete(id)

	log.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
