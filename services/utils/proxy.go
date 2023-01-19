package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Redirect from given path to the same path but on the another address
func HandleWithProxy(path string, address string) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(address)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})
}
