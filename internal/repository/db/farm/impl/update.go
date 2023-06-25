package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) UpdateFarm(ctx context.Context, req model.UpsertFarm) error {
	var (
		err error
	)

	_, err = r.db.ExecContext(ctx, queryUpdateFarm, req.FarmID, req.Description)
	if err != nil {
		log.Printf("[repository.db.farm.UpdateFarm] err: %v", err)
		return err
	}

	return err
}
