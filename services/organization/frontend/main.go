package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/organization/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/organization/organizations/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/organizations.html")
	})

	http.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir("../../../share"))))
	fmt.Println("Organizations service started...")
	http.ListenAndServe(":8082", nil)
}
