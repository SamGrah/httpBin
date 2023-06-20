package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"bin-manager/internal/db-service"
)

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

func LogRequestToBin(binId string, requestDetails *db_service.HttpRequestContents) (error) {
	binInDb, err := db_service.BinIdExists(binId)
	if err != nil {
		log.Fatal(err)
	}

	if !binInDb {
		err = errors.New("http request to be logged in unknown bin")
		log.Fatal(err)
		return err 
	}
	
	err = db_service.AddRequestToBin(binId, *requestDetails)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func FetchRequestsFromBin(binId string) (*[]db_service.HttpRequestContents, error) {
	binContents, err := db_service.GetBinContents(binId)
	if err != nil {
		log.Fatal(err)
		emptySlice := make([]db_service.HttpRequestContents, 0)
		return &emptySlice, err
	}

	return binContents, nil
}
