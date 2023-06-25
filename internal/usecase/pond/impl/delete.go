package impl

import (
	"aquafarm/common/model"
	"context"
)

func (u *usecase) DeletePondByID(ctx context.Context, req model.DeletePondByID) error {
	return u.rDbPond.DeletPondByID(ctx, req)
}
