package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/customer/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/customer/customers/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/customers.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	fmt.Println("Customers service started...")
	http.ListenAndServe(":8085", nil)
}
