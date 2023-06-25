package pond

import (
	"aquafarm/common/model"
	"context"
)

type Usecase interface {
	InsertPond(ctx context.Context, req model.InsertPondReq) error
	GetPond(ctx context.Context, req model.GetPondReq) ([]model.Pond, error)
	GetPondByID(ctx context.Context, req model.GetPondByID) (model.Pond, error)
	DeletePondByID(ctx context.Context, req model.DeletePondByID) error
	UpsertPond(ctx context.Context, req model.UpsertPond) error
}
