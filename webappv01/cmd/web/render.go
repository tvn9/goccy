package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var pathToTemplates = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	// User *data.User
}

func (c *Config) render(w http.ResponseWriter, r *http.Request, t string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.html", pathToTemplates),
		fmt.Sprintf("%s/header.partial.html", pathToTemplates),
		fmt.Sprintf("%s/navbar.partial.html", pathToTemplates),
		fmt.Sprintf("%s/footer.partial.html", pathToTemplates),
		fmt.Sprintf("%s/alerts.partial.html", pathToTemplates),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplates, t))

	templateSlice = append(templateSlice, partials...)

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		c.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, c.AddDefaultData(td, r)); err != nil {
		c.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Config) AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = c.Session.PopString(r.Context(), "flash")
	td.Warning = c.Session.PopString(r.Context(), "warning")
	td.Error = c.Session.PopString(r.Context(), "error")
	if c.IsAuthenticated(r) {
		td.Authenticated = true
		// TODO - get more user information
	}
	td.Now = time.Now()
	return td
}

func (c *Config) IsAuthenticated(r *http.Request) bool {
	return c.Session.Exists(r.Context(), "userID")
}
