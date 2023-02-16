package db_service

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type BinId string
type HttpRequestContents map[string]string
type BinContents []HttpRequestContents

type CollectionDetails struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type HttpRequestDetails struct {
	BinId   
	HttpRequestContents
}

type Bin struct {
	BinId
	*BinContents 
}
