package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

// function to insert new pond
func (r *repository) InsertPond(ctx context.Context, data model.InsertPondReq) error {
	var (
		err error
	)

	// insert data to database
	_, err = r.db.ExecContext(ctx, queryInsertPond, data.FarmID, data.PondName, data.Description)
	if err != nil {
		log.Printf("[repository.db.pond.InsertFarm] err: %v", err)
		return err
	}

	return err
}
