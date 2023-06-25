package impl

import (
	"aquafarm/common/model"
	"context"
	"log"
)

func (r *repository) GetFarmByName(ctx context.Context, req model.GetFarmByNameReq) ([]model.GetFarmByNameRes, error) {
	var (
		result []model.GetFarmByNameRes
		err    error
	)

	err = r.db.SelectContext(ctx, &result, queryGetFarmByName, req.FarmName)
	if err != nil {
		log.Printf("[repository.db.farm.GetFarmByName] err: %v", err)
		return result, err
	}

	return result, err
}

func (r *repository) GetFarm(ctx context.Context, filter model.GetFarmReq) ([]model.GetFarmResult, error) {
	var (
		result []model.GetFarmResult
		err    error
	)

	limit := filter.PerPage
	offset := filter.PerPage * (filter.Page - 1)

	err = r.db.SelectContext(ctx, &result, queryGetFarm, limit, offset)
	if err != nil {
		log.Printf("[repository.db.farm.GetFarm] err: %v", err)
		return result, err
	}

	return result, err
}

func (r *repository) GetFarmByID(ctx context.Context, req model.GetFarmByIDReq) (model.GetFarmByIDRes, error) {
	var (
		result model.GetFarmByIDRes
		err    error
	)

	err = r.db.GetContext(ctx, &result, queryGetFarmByID, req.FarmID)
	if err != nil {
		log.Printf("[repository.db.farm.GetFarmByID] err: %v", err)
		return result, err
	}

	return result, err
}
