package impl

import (
	"aquafarm/common/model"
	"context"
)

func (u *usecase) UpsertFarm(ctx context.Context, req model.UpsertFarm) error {
	var (
		err error
	)

	// get farm by name
	farmResult, err := u.rDbFarm.GetFarmByName(ctx, model.GetFarmByNameReq{
		FarmName: req.FarmName,
	})
	if err != nil {
		return err
	}

	// if there is no data we will insert new data
	if len(farmResult) == 0 {
		err = u.rDbFarm.InsertFarm(ctx, model.InsertFarmReq{
			FarmName:    req.FarmName,
			Description: req.Description,
		})
		return err
	}

	// assign farm id
	req.FarmID = farmResult[0].FarmID
	// update existing data
	err = u.rDbFarm.UpdateFarm(ctx, req)

	return err
}
