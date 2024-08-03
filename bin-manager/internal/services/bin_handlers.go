package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	db_service "bin-manager/internal/db-service"
)

func CreateNewBin() (string, error) {
	var binId string
	for {
		binId = generateNewBinId()

		duplicateBinId, err := db_service.BinIdExists(binId)
		if err != nil {
			log.Printf("Failed to verify binId %s", binId)
			return "", err
		}

		if !duplicateBinId {
			break
		}
	}

	err := db_service.CreateNewBin(binId)
	if err != nil {
		fmt.Printf("Failed to generate new bin %s", binId)
		return "", err
	}

	return binId, nil
}

func generateNewBinId() string {
	rand.Seed(time.Now().UnixNano())
	chars := "abcdefghijklmnopqrstuvwxyz123456789"
	min := 0
	max := len(chars) - 1

	binId := ""
	for idx := 0; idx < 6; idx++ {
		index := rand.Intn(max - min + 1)
		character := string(chars[index])
		binId += character
	}

	return binId
}

func LogRequestToBin(binId string, requestDetails *db_service.HttpRequest) error {
	binInDb, err := db_service.BinIdExists(binId)
	if err != nil {
		log.Printf("Failed to verify binId %s", binId)
	}

	if !binInDb {
		err = errors.New("http request to be logged in unknown bin")
		log.Printf("Failed to log request to non-existant bin %s", binId)
		return err
	}

	err = db_service.AddRequestToBin(binId, *requestDetails)
	if err != nil {
		log.Printf("Failed to log request to bin %s", binId)
		return err
	}

	return nil
}

func FetchRequestsFromBin(binId string) (*[]db_service.HttpRequest, error) {
	binInDb, err := db_service.BinIdExists(binId)
	if err != nil {
		log.Printf("Failed to verify binId %s", binId)
		return nil, err
	}

	if !binInDb {
		err = errors.New("http request to be logged in unknown bin")
		log.Printf("Failed to log request to non-existant bin %s", binId)
		return nil, err
	}

	binContents, err := db_service.GetBinContents(binId)
	if err != nil {
		log.Printf("Failed to fetch bin contents for bin %s", binId)
		emptySlice := make([]db_service.HttpRequest, 0)
		return &emptySlice, err
	}

	return binContents, nil
}
