package main

import (
        "net/http"
        "html/template"
    )

type Page struct {
    Name string
}

var templates = template.Must( template.ParseFiles("./templates/index.html"))

func Index(w http.ResponseWriter, r *http.Request) {
    page := Page{Name:"qamt"}
    if name := r.FormValue("name"); name != "" {
        page.Name = name
    }

    if err := templates.ExecuteTemplate(w, "index.html", page); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
