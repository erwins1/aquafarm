package impl

import (
	rDbFarm "aquafarm/internal/repository/db/farm"
	uFarm "aquafarm/internal/usecase/farm"
)

type usecase struct {
	rDbFarm rDbFarm.Repository
}

func New(rDbFarm rDbFarm.Repository) uFarm.Usecase {
	return &usecase{
		rDbFarm: rDbFarm,
	}
}
