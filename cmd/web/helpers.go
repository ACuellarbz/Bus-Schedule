package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
)

// cache will hold the rendered templates so
// that we do not have to go to disk every time
var cache = map[string]*template.Template{}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *templateData) {
	// Check if the cache has the template for the passed in
	// template string. If it does not, then we call the
	// createTemplateCache() function to parse the file and
	// place it into the cache
	_, ok := cache[tmpl]
	if !ok {
		err := createTemplateCache(tmpl)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	}
	// The template is in the cache, so read it and
	// execute it
	ts := cache[tmpl]
	err := ts.Execute(w, data)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// The createTemplateCache() function accepts a string
// representing a template. The template is then parsed
// and place into a template cache. An error is returned
func createTemplateCache(tmpl string) error {
	templates := []string{
		fmt.Sprintf("./internal/models/ui/html/%s", tmpl),
		"./internal/models/ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// Add the template to the map
	cache[tmpl] = ts
	return nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	// deal with the error status
	http.Error(w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionsManager.Exists(r.Context(), "authenticatedUserID")
}
