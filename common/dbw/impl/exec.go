package impl

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (r *repository) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = r.dbConn.Rebind(query)

	rows, err := r.dbConn.ExecContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return rows, err
	}

	return rows, nil
}
