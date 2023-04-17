package models

type HttpRequests struct {
	LoggedRequests []string
}

type Bin struct {
	BinId string
	Contents *HttpRequests 
}

type BinMgmtRepository interface {
	CreateNewBin() (*Bin, error)
}