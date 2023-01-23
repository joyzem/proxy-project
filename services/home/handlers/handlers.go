package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/index.html")
}

func GolangHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../static/html/about-golang.html")
}

func LabsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	labName := fmt.Sprintf("lab%s.md", id)
	title := fmt.Sprintf("Лабораторная работа %s", id)
	labTemplate, err := template.ParseFiles("../static/html/lab.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	type LabData struct {
		Title   string
		LabName string
	}
	data := LabData{
		Title:   title,
		LabName: labName,
	}
	if err := labTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
