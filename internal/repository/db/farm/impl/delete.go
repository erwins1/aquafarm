package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) DeleteFarmByID(ctx context.Context, req model.DeleteFarmByIDReq) error {
	var (
		err error
	)

	_, err = r.db.ExecContext(ctx, queryDeleteFarm, req.FarmID)
	if err != nil {
		log.Printf("[repository.db.farm.DeleteFarmByID] err: %v", err)
		return err
	}

	return err
}
