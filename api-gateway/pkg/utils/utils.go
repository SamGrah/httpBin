package utils

import (
	"fmt"
	"log"
)

func LogError(msg string, err error) {
	errorMsg := fmt.Sprintf("Error: %s", err)
	log.Printf("%s: %s", msg, errorMsg)
}
