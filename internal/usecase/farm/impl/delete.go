package impl

import (
	"aquafarm/common/model"
	"context"
)

func (u *usecase) DeleteFarmByID(ctx context.Context, req model.DeleteFarmByIDReq) error {
	return u.rDbFarm.DeleteFarmByID(ctx, req)
}
