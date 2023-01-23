package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

// Redirect from given path to the same path but on the another address
func HandleWithProxy(router *mux.Router, path string, address string) {
	router.PathPrefix(path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, _ := url.Parse(address)
		proxy := httputil.NewSingleHostReverseProxy(url)
		proxy.ServeHTTP(w, r)
	})
}
