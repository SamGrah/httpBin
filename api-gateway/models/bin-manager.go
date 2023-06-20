package models

type BinId struct {
	BinId string
}

type HttpRequest struct {
	Contents map[string]string
}

type Contents struct {
	LoggedRequests []*HttpRequest
}

type Bin struct {
	BinId
	Contents []*HttpRequest	
}

type BinMgmtRepository interface {
	CreateNewBin() (*Bin, error)
}