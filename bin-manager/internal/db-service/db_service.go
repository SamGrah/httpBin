package db_service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"

	"api-gateway/pkg/utils"
)

const dbPort = "27017"
const dbIP = "mongo_db"

func Init() error {
	fmt.Printf("Connecting to MongoDB...")
	dbClient, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		utils.LogError("db error", err)
		return err
	}

	expireIndex, err := expireIndexExists(dbClient.Collection)
	if err != nil {
		utils.LogError("db error", err)
		log.Fatal(err)
		return err
	}

	if !expireIndex {
		err = createExpirationIndex(dbClient.Collection)
		if err != nil {
			utils.LogError("failed to create expiration index", err)
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
		utils.LogError("failed to connected to MongoDB", err)
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
		fmt.Println("Error encountered when disconnecting from MongoDB!")
		log.Fatal(err)
	}
	fmt.Println("Disconnected client from MongoDB!")
}

func getDbCollectionDetails(DbName string, CollectionName string) (*CollectionDetails, error) {
	client, err := createDbConn()
	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)
	collectionDetails := CollectionDetails{client, collection}
	return &collectionDetails, nil
}

func expireIndexExists(binsCollection *mongo.Collection) (bool, error) {
	cursor, err := binsCollection.Indexes().List(context.TODO())
	if err != nil {
		return false, errors.Join(
			errors.New("failed to retrieve list of indexes for bins collection"),
			err,
		)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		defer cursor.Close(context.TODO())
		return false, errors.Join(
			errors.New("failed to decode list of indexes for bins collection"),
			err,
		)
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
			Key:   "created_at",
			Value: bsonx.Int32(1),
		}},
		Options: expirationDuration,
	}
	createIndexResult, err := binsCollection.Indexes().CreateOne(context.Background(), mod)
	if err != nil {
		return errors.Join(
			fmt.Errorf("returned result: %s", createIndexResult),
			fmt.Errorf("failed to create index: %+v", err),
		)
	}

	log.Println("Created document expiration index for bins collection")
	return nil
}

func CreateNewBin(binId string) error {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		return errors.Join(fmt.Errorf("failed to fetch Db collection"), err)
	}

	emptyHttpRequestsSlice := make([]HttpRequest, 0)

	newBin := &Bin{
		BinId:    binId,
		Requests: emptyHttpRequestsSlice,
	}
	insertResult, err := connDetails.Collection.InsertOne(context.Background(), newBin)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to insert one record"), err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	closeDbConn(connDetails.Client)
	return nil
}

func BinIdExists(binId string) (bool, error) {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		return false, errors.Join(fmt.Errorf("failed to fetch Db collection"), err)
	}

	var result Bin
	var binExists bool

	filter := bson.D{primitive.E{Key: "binid", Value: binId}}
	err = connDetails.Collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			binExists = false
			err = nil
		} else {
			err = errors.Join(fmt.Errorf("failed to find and decode bin"), err)
		}
	} else {
		binExists = true
	}

	closeDbConn(connDetails.Client)
	return binExists, err
}

func AddRequestToBin(binId string, request HttpRequest) error {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		return errors.Join(fmt.Errorf("failed to fetch Db collection"), err)
	}

	filter := bson.D{{Key: "binid", Value: binId}}
	update := bson.M{
		"$push": bson.M{
			"requests": bson.M{
				"hostIp":   request.HostIp,
				"recieved": request.Recieved,
				"contents": request.Contents,
			},
		},
	}

	_, err = connDetails.Collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.Join(fmt.Errorf("failed to update one record"), err)
	}

	closeDbConn(connDetails.Client)
	return nil
}

func GetBinContents(binId string) (*[]HttpRequest, error) {
	connDetails, err := getDbCollectionDetails("httpBin", "bins")
	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to fetch Db collection"), err)
	}

	if err != nil {
		return nil, errors.Join(fmt.Errorf("failed to fetch Db collection"), err)
	}

	var bin Bin
	filter := bson.D{primitive.E{Key: "binid", Value: binId}}

	err = connDetails.Collection.FindOne(context.TODO(), filter).Decode(&bin)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			bin = Bin{}
		} else {
			err = errors.Join(fmt.Errorf("failed to find and decode bin"), err)
			return nil, err
		}
	}

	log.Printf("Fetched bin %s contents: %+v", bin.BinId, bin.Requests)
	closeDbConn(connDetails.Client)
	return &bin.Requests, nil
}
