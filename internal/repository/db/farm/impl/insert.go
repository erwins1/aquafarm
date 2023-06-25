package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) InsertFarm(ctx context.Context, data model.InsertFarmReq) error {
	var (
		err error
	)

	_, err = r.db.ExecContext(ctx, queryInsertFarm, data.FarmName)
	if err != nil {
		log.Printf("[repository.db.farm.InsertFarm] err: %v", err)
		return err
	}

	return err
}
