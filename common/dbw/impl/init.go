package impl

import (
	"github.com/jmoiron/sqlx"

	// repo
	rDbw "aquafarm/common/dbw"
)

type repository struct {
	dbConn *sqlx.DB
}

func New(
	dbConn *sqlx.DB,
) rDbw.Repository {
	return &repository{
		dbConn: dbConn,
	}
}
