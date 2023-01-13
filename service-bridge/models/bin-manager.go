package models

type HttpRequests struct {
	LoggedRequests []string
}

type Bin struct {
	Id string
	Contents *HttpRequests 
}

type BinMgmtRepository interface {
	CreateNewBin() (*Bin, error)
}