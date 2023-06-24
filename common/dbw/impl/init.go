package impl

import (
	"github.com/jmoiron/sqlx"

	// repo
	rDbw "aquafarm/common/dbw"
)

type repository struct {
	dbSlave *sqlx.DB
}

func New(
	dbSlave *sqlx.DB,
) rDbw.Repository {
	return &repository{
		dbSlave: dbSlave,
	}
}
