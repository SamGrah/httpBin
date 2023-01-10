package db_service

import (
	"context"
	"log"
	"fmt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bins_collection *mongo.Collection

type HttpRequestDetails struct {
	Content string
}

type Bin struct {
	BinId string
	RequestContent HttpRequestDetails	
}

func CreateDbConn(dbIP string, dbPort string) (*mongo.Client, error) {
	// Set client options
	connectionString := fmt.Sprintf("mongodb://admin:password@%s:%s/", dbIP, dbPort)
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed MongoDB connection")
		log.Fatal(err)
		return nil, err 
	}
	fmt.Println("Connected to MongoDB!")
	// Close the db connection when bin_manager process ends
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Disconnected client from MongoDB!")
	// }()

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed MongoDB ping")
		log.Fatal(err)
		return nil, err 
	}
	fmt.Println("MongoDB database successfully pinged")

	// bins_collection = client.Database("bins").Collection("bins")
	return client, nil
}

// func FindBin(binId string) (string, error) {
// 	newBin := Bin{"newID", HttpRequestDetails{"http request header & payload"}}
// 	insertResult, err := bins_collection.InsertOne(context.TODO(), newBin)
// 	if err != nil {
// 			fmt.Println("Failed to insert one record")
// 			log.Fatal(err)
// 	}
// 	fmt.Println("Inserted a single document: ", insertResult.InsertedID)	

// }

// func CreateNewBin() (*mongo.Collection, error) {}

// func AddRequestToBin() (*mongo.Collection, error) {}