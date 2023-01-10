package db_service

import (
	"context"
	"log"
	"fmt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbPort = "27017"
const dbIP = "mongo_db"

type CollectionDetails struct {
	Client *mongo.Client
	Collection *mongo.Collection
}

type HttpRequestDetails struct {
	Content string
}

type Bin struct {
	BinId string
	RequestContent HttpRequestDetails	
}

func createDbConn() (*mongo.Client, error) {
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

func closeDbConn(client *mongo.Client) {
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
		fmt.Println("Error encountered when disconnecting from MongoDB!")
	}
	fmt.Println("Disconnected client from MongoDB!")
}

func getDbCollectionDetails(DbName string, CollectionName string) (*CollectionDetails, error) {
	client, err := createDbConn()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	collectionDetails := CollectionDetails{ client, collection } 
	return &collectionDetails, nil
}

func CreateNewBin(binId string) error {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
			fmt.Println("Failed to fetch Db collection")
			log.Fatal(err)
			return err
	}

	newBin := Bin{binId, HttpRequestDetails{"http request header & payload"}}
	insertResult, err := connDetails.Collection.InsertOne(context.Background(), newBin)
	if err != nil {
			fmt.Println("Failed to insert one record")
			log.Fatal(err)
			return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	closeDbConn(connDetails.Client)
	return nil 
}

// func FindBin(binId string) (string, error) {}
// func AddRequestToBin() (*mongo.Collection, error) {}