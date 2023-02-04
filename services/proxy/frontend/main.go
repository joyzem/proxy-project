package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/proxy", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/proxy/proxies", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/proxies.html")
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../static"))))
	fmt.Println("Listening on 8086...")
	http.ListenAndServe(":8086", nil)
}
