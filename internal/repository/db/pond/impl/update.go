package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) UpdatePond(ctx context.Context, req model.UpsertPond) error {
	var (
		err error
	)

	_, err = r.db.ExecContext(ctx, queryUpdatePond, req.PondID, req.FarmID, req.Description)
	if err != nil {
		log.Printf("[repository.db.pond.UpdatePond] err: %v", err)
		return err
	}

	return err
}
