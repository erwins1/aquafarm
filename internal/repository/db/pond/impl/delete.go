package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) DeletPondByID(ctx context.Context, req model.DeletePondByID) error {
	var (
		err error
	)

	_, err = r.db.ExecContext(ctx, queryDeletePond, req.PondID)
	if err != nil {
		log.Printf("[repository.db.pond.DeletPondByID] err: %v", err)
		return err
	}

	return err
}
