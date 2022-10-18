package main

import (
	"bin-manager/generated/adapters"
	"context"
)

func (s *binManagerServer) GenerateNewBin(ctx context.Context, params *binManager.Params) (*binManager.BinId, error) {
	binId := binManager.BinId{
		BinId: "xv7ty",
	}
	return &binId, nil 
}
