package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/account/frontend/router"
	"github.com/joyzem/proxy-project/services/base"
)

func main() {
	router := router.GetRouter()

	fmt.Println("Listening on 8083...")
	if err := http.ListenAndServe(":8083", router); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
