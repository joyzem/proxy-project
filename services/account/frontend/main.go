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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	fmt.Println("Listening on 8083...")
	http.ListenAndServe(":8083", nil)
}
