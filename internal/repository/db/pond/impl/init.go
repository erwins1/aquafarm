package impl

import (
	"aquafarm/common/dbw"
	rDbPond "aquafarm/internal/repository/db/pond"
)

type repository struct {
	db dbw.Repository
}

func New(db dbw.Repository) rDbPond.Repository {
	return &repository{
		db: db,
	}
}
