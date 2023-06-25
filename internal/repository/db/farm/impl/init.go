package impl

import (
	"aquafarm/common/dbw"
	rDbFarm "aquafarm/internal/repository/db/farm"
)

type repository struct {
	db dbw.Repository
}

func New(db dbw.Repository) rDbFarm.Repository {
	return &repository{
		db: db,
	}
}
