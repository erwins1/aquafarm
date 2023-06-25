package farm

import (
	"aquafarm/common/model"
	"context"
)

type Usecase interface {
	InsertFarm(ctx context.Context, req model.InsertFarmReq) error
	GetFarm(ctx context.Context, req model.GetFarmReq) ([]model.GetFarmResult, error)
	GetFarmByID(ctx context.Context, req model.GetFarmByIDReq) (model.GetFarmByIDRes, error)
}
