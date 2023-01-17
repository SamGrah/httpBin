package db_service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CollectionDetails struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type HttpRequestDetails struct {
	BinId	string
	Content string
}

type Bin struct {
	BinId    string
	Requests []map[string]string 
}
