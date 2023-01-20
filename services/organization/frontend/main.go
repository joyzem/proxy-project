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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	fmt.Println("Listening on 8082...")
	http.ListenAndServe(":8082", nil)
}
