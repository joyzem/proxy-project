package utils

import (
	"fmt"
)

func LogError(err error) {
	fmt.Printf("error: %s\n", err.Error())
}
