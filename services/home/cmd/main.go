package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/home/router"
	"github.com/joyzem/proxy-project/services/utils"
)

func main() {

	hanlder := router.GetRouter()

	fmt.Println("Listening on 80...")
	if err := http.ListenAndServe(":80", hanlder); err != nil {
		utils.LogError(err)
		os.Exit(-1)
	}
}
