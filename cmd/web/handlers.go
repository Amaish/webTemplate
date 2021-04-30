package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"gitlab.com/Amaish/webTemplate/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.blogs.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	for _, blog := range s {
		fmt.Fprintf(w, "%v\n\n", blog)
	}
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/sidebar.partial.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showBlog(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	s, err := app.blogs.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}
	files := []string{
		"./ui/html/show.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/sidebar.partial.html",
		"./ui/html/footer.partial.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = ts.Execute(w, s)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) newBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := "Pumwani Maternity Hospital crisis"
	author := "Madam Muthoni"
	content := `In a recent press conference held on 5 October at Silver Springs Hotel, 
	Nairobi, to highlight the gaps in information about the crisis at Pumwani Maternity Hospital.`
	expire := "7"
	id, err := app.blogs.Insert(title, author, content, expire)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/blogs?id=%d", id), http.StatusSeeOther)
}
