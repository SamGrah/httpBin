package db_service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type HttpRequestContents map[string]string

type CollectionDetails struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type HttpRequestDetails struct {
	BinId string
	HttpRequestContents
}

type Bin struct {
	BinId   string                `bson:"binid"`
	Requets []HttpRequestContents `bson:"requests"`
}
