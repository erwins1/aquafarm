package impl

import (
	"aquafarm/common/model"
	"context"
	"errors"
)

func (u *usecase) UpsertPond(ctx context.Context, req model.UpsertPond) error {
	var (
		err error
	)

	// check farm id exist or not
	farmResult, err := u.rDbFarm.GetFarmByID(ctx, model.GetFarmByIDReq{
		FarmID: req.FarmID,
	})
	if err != nil {
		return err
	}
	if farmResult.FarmID == 0 {
		return errors.New("invalid farm id")
	}

	// if there is no data we will insert new data
	pondResult, err := u.rDbPond.GetPondByName(ctx, model.GetPondByNameReq{
		PondName: req.PondName,
	})
	if err != nil {
		return err
	}
	if len(pondResult) == 0 {
		err = u.rDbPond.InsertPond(ctx, model.InsertPondReq{
			FarmID:      req.FarmID,
			PondName:    req.PondName,
			Description: req.Description,
		})
		return err
	}

	// assign farm id
	req.PondID = pondResult[0].PondID
	// update existing data
	err = u.rDbPond.UpdatePond(ctx, req)

	return err
}
