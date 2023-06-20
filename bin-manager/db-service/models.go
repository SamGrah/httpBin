package db_service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type HttpRequest struct {
	HostIp string
	Recieved *timestamppb.Timestamp
	Contents map[string]string
}

type CollectionDetails struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

type HttpRequestDetails struct {
	BinId string
	HttpRequest
}

type Bin struct {
	BinId    string        `bson:"binid"`
	Requests []HttpRequest `bson:"requests"`
}
