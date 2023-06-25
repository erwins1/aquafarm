package farm

import (
	"aquafarm/common/model"
	"context"
)

type Repository interface {
	InsertFarm(ctx context.Context, data model.InsertFarmReq) error
	GetFarmByName(ctx context.Context, req model.GetFarmByNameReq) ([]model.GetFarmByNameRes, error)
	GetFarmByID(ctx context.Context, req model.GetFarmByIDReq) (model.GetFarmByIDRes, error)
	GetFarm(ctx context.Context, filter model.GetFarmReq) ([]model.GetFarmResult, error)
	DeleteFarmByID(ctx context.Context, req model.DeleteFarmByIDReq) error
	UpdateFarm(ctx context.Context, req model.UpsertFarm) error
}
