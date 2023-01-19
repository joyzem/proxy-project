package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/employee/employees/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/employees.html")
	})
	fmt.Println("Employees service started...")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	http.Handle("/employee/js/", http.StripPrefix("/employee/js/", http.FileServer(http.Dir("views"))))
	http.ListenAndServe(":8084", nil)
}
