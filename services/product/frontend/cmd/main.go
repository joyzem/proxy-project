package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/product/frontend/router"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
