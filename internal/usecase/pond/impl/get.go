package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"context"
	"errors"
)

// usecase to get pond
func (u *usecase) GetPond(ctx context.Context, req model.GetPondReq) ([]model.Pond, error) {
	var (
		result []model.Pond
		err    error
	)

	// get data from repo
	result, err = u.rDbPond.GetPond(ctx, req)

	return result, err
}

func (u *usecase) GetPondByID(ctx context.Context, req model.GetPondByID) (model.Pond, error) {
	var (
		result model.Pond
		err    error
	)

	result, err = u.rDbPond.GetPondByID(ctx, req)
	if err != nil {
		return result, err
	}
	if result.PondID == 0 {
		return result, errors.New(constant.ErrNotFound)
	}

	return result, err
}
