package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

// get pond based on farm id
func (r *repository) GetPondByFarmID(ctx context.Context, req model.GetPondByFarmIDReq) ([]model.Pond, error) {
	var (
		result = []model.Pond{}
		err    error
	)
	// get data from db
	err = r.db.SelectContext(ctx, &result, queryGetPondByFarmID, req.FarmID)
	if err != nil {
		log.Printf("[repository.db.pond.GetPondByFarmID] err: %v", err)
		return result, err
	}

	return result, err
}

// get pond by name
func (r *repository) GetPondByName(ctx context.Context, req model.GetPondByNameReq) ([]model.GetPondByNameRes, error) {
	var (
		result = []model.GetPondByNameRes{}
		err    error
	)
	// get data from db
	err = r.db.SelectContext(ctx, &result, queryGetPondByName, req.PondName)
	if err != nil {
		log.Printf("[repository.db.pond.GetPondByName] err: %v", err)
		return result, err
	}

	return result, err
}

// get pond func
func (r *repository) GetPond(ctx context.Context, filter model.GetPondReq) ([]model.Pond, error) {
	var (
		result = []model.Pond{}
		err    error
	)

	// calculate limit and offset
	limit := filter.PerPage
	offset := filter.PerPage * (filter.Page - 1)

	// get pond data
	err = r.db.SelectContext(ctx, &result, queryGetPond, limit, offset)
	if err != nil {
		log.Printf("[repository.db.pond.GetPond] err: %v", err)
		return result, err
	}

	return result, err
}

// get pond by pond id
func (r *repository) GetPondByID(ctx context.Context, req model.GetPondByID) (model.Pond, error) {
	var (
		result = model.Pond{}
		err    error
	)

	err = r.db.GetContext(ctx, &result, queryGetPondByID, req.PondID)
	if err != nil {
		log.Printf("[repository.db.pond.GetPondByID] err: %v", err)
		return result, err
	}

	return result, err
}
