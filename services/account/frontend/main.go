package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/account/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/account/accounts/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/accounts.html")
	})
	http.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir("../../../share"))))
	fmt.Println("Accounts service started...")
	http.ListenAndServe(":8083", nil)
}
