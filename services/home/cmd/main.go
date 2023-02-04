package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joyzem/proxy-project/services/base"
	"github.com/joyzem/proxy-project/services/home/router"
)

func main() {

	hanlder := router.GetRouter()

	fmt.Println("Listening on 80...")
	if err := http.ListenAndServe(":80", hanlder); err != nil {
		base.LogError(err)
		os.Exit(-1)
	}
}
