package pond

import (
	"aquafarm/common/model"
	"context"
)

type Repository interface {
	GetPondByFarmID(ctx context.Context, req model.GetPondByFarmIDReq) ([]model.Pond, error)
	GetPondByName(ctx context.Context, req model.GetPondByNameReq) ([]model.GetPondByNameRes, error)
	InsertPond(ctx context.Context, data model.InsertPondReq) error
	GetPond(ctx context.Context, filter model.GetPondReq) ([]model.Pond, error)
	GetPondByID(ctx context.Context, req model.GetPondByID) (model.Pond, error)
	DeletPondByID(ctx context.Context, req model.DeletePondByID) error
}
