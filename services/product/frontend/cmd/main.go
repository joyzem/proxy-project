package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/product/frontend/router"
	"github.com/joyzem/proxy-project/services/utils"
)

func main() {

	handler := router.GetRouter()

	fmt.Println("Listening on 8081...")
	if err := http.ListenAndServe(":8081", handler); err != nil {
		utils.LogError(err)
		os.Exit(-1)
	}
}
