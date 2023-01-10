package services 

import (
	"fmt"
	"log"

	"bin-manager/db-service"
)

func generateNewBinId() string {
	binId := "xv7ty"
	return binId
}

func GenerateNewBin() (string, error) {
	binId := generateNewBinId() 
	err := db_service.CreateNewBin(binId)
	if err != nil {
		fmt.Println("Failed to generate new bin")
		log.Fatal(err)
		return "", err
	}

	return binId, nil
}
