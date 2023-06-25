package impl

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (r *repository) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.dbConn.Rebind(query)

	rows, err := r.dbConn.QueryContext(ctx, query, args)
	if err != nil && err != sql.ErrNoRows {
		return rows, err
	}

	return rows, nil
}
