package impl

import (
	"aquafarm/common/model"
	"context"
	"errors"
)

func (u *usecase) InsertFarm(ctx context.Context, req model.InsertFarmReq) error {
	var (
		err error
	)
	farmResult, err := u.rDbFarm.GetFarmByName(ctx, model.GetFarmByNameReq{
		FarmName: req.FarmName,
	})
	if err != nil {
		return err
	}

	if len(farmResult) > 0 {
		return errors.New("duplicate farm name")
	}

	err = u.rDbFarm.InsertFarm(ctx, req)

	return err
}
