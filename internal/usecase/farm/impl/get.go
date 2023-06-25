package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"context"
	"errors"
)

func (u *usecase) GetFarm(ctx context.Context, req model.GetFarmReq) ([]model.GetFarmResult, error) {
	var (
		result []model.GetFarmResult
		err    error
	)

	result, err = u.rDbFarm.GetFarm(ctx, req)

	return result, err
}

func (u *usecase) GetFarmByID(ctx context.Context, req model.GetFarmByIDReq) (model.GetFarmByIDRes, error) {
	var (
		result model.GetFarmByIDRes
		err    error
	)

	result, err = u.rDbFarm.GetFarmByID(ctx, req)
	if err != nil {
		return result, err
	}
	if result.FarmID == 0 {
		return result, errors.New(constant.ErrNotFound)
	}

	result.Ponds, err = u.rDbPond.GetPondByFarmID(ctx, model.GetPondByFarmIDReq{
		FarmID: result.FarmID,
	})
	if err != nil {
		return result, err
	}

	return result, err
}
