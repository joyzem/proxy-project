package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/proxy/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/index.html")
	})
	http.HandleFunc("/proxy/proxies/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/proxies.html")
	})
	http.Handle("/share/", http.StripPrefix("/share/", http.FileServer(http.Dir("../../../share"))))
	fmt.Println("Proxies service started...")
	http.ListenAndServe(":8086", nil)
}
