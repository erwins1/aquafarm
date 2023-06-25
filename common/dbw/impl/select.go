package impl

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func (r *repository) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	query = r.dbConn.Rebind(query)

	err = r.dbConn.SelectContext(ctx, dest, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}

func (r *repository) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {

	query, args, err := sqlx.In(query, args...)
	if err != nil {
		return err
	}
	query = r.dbConn.Rebind(query)

	err = r.dbConn.GetContext(ctx, dest, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
