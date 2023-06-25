package impl

import (
	"aquafarm/common/model"
	"context"
	"errors"
)

func (u *usecase) InsertPond(ctx context.Context, req model.InsertPondReq) error {
	var (
		err error
	)

	// check duplicate pond name
	pondResult, err := u.rDbPond.GetPondByName(ctx, model.GetPondByNameReq{
		PondName: req.PondName,
	})
	if err != nil {
		return err
	}
	if len(pondResult) > 0 {
		return errors.New("duplicate pond name")
	}

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

	// insert pond
	err = u.rDbPond.InsertPond(ctx, req)

	return err
}
