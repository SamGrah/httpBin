package db_service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const dbPort = "27017"
const dbIP = "mongo_db"


func Init() error {
	fmt.Printf("Connecting to MongoDB...")
	dbClient, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		log.Fatal(err)
		return err
	}

	expireIndex, err := expireIndexExists(dbClient.Collection)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if !expireIndex {
		err = createExpirationIndex(dbClient.Collection)
		if err != nil {
			fmt.Println("Failed to create expiration index")
			log.Fatal(err)
			return err
		}
	}

	closeDbConn(dbClient.Client)
	return nil
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

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Failed MongoDB ping")
		log.Fatal(err)
		return nil, err 
	}
	fmt.Println("MongoDB database successfully pinged")

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

func expireIndexExists(binsCollection *mongo.Collection) (bool, error) {
	cursor, err := binsCollection.Indexes().List(context.TODO())
	if err != nil {
			fmt.Println("Failed to retrieve list of indexes for bins collection")
			log.Fatal(err)
			return false, err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)	
	}
	for _, result := range results {
		if result["expireAfterSeconds"] != nil {
			return true, nil
		}
	}

	return false, nil
}

func createExpirationIndex(binsCollection *mongo.Collection) error {
	expirationDuration := options.Index().SetExpireAfterSeconds(60 * 60 * 24)
	mod := mongo.IndexModel{
		Keys: bsonx.Doc{{
			Key: "created_at",
			Value: bsonx.Int32(1), 
		}}, 
		Options: expirationDuration,
	}
	createIndexResult, err := binsCollection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		fmt.Println("returned result: ", createIndexResult)
		fmt.Println("failed to create index: ", err)
		return err
	}

	fmt.Println("Created document expiration index for bins collection")
	return nil
}

func CreateNewBin(binId string) error {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
			fmt.Println("Failed to fetch Db collection")
			log.Fatal(err)
			return err
	}

	// var emptyHttpRequestsSlice []HttpRequestDetails
	emptyHttpRequestsSlice := make([]HttpRequestDetails, 0)
	newBin := Bin{
		BinId: binId, 
		Requests: emptyHttpRequestsSlice,
	}
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

func BinIdExists(binId string) (bool, error) {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
			fmt.Println("Failed to fetch Db collection")
			log.Fatal(err)
			return false, err
	}

	var result Bin
	var binExists bool

	filter := bson.D{primitive.E{Key: "binid", Value: binId}}
	err = connDetails.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			binExists = false	
		} else {
			log.Fatal(err)
		}	
	} else {
		binExists = true
	}

	closeDbConn(connDetails.Client)
	return binExists, nil
}

func AddRequestToBin(binId string, content map[string]string) error {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
			fmt.Println("Failed to fetch Db collection")
			log.Fatal(err)
			return err
	}

	filter := bson.D{{Key:"binid", Value: binId}}		
	update := bson.M{ 
		"$push": bson.M{ "requests": content },
	}

	_, err = connDetails.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("error adding request contents to bin: %s", binId)
		log.Fatal(err)
		return err
	}

	closeDbConn(connDetails.Client)
	return nil
}