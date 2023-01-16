package services 

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"bin-manager/db-service"
)

func generateNewBinId() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz123456789"
	min := 0 
	max := len(chars) - 1

	binId := ""
	for i := 0; i < 6; i++ {
		index := rand.Intn(max - min + 1)
		character := string(chars[index])
		binId += character
	}

	return binId
}

func CreateNewBin() (string, error) {
	var binId string
	for {
		binId = generateNewBinId() 

		duplicateBinId, err := db_service.BinIdExists(binId)
		if err != nil {
			log.Fatal(err)
		}

		if !duplicateBinId {
			break
		}
	}

	err := db_service.CreateNewBin(binId)
	if err != nil {
		fmt.Println("Failed to generate new bin")
		log.Fatal(err)
		return "", err
	}

	return binId, nil
}

func LogRequestToBin(binId string, requestDetails map[string]string) (error) {
	err := db_service.AddRequestToBin(binId, requestDetails)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
