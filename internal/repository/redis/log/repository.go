package log

import (
	"aquafarm/common/model"
	"context"
)

type Repository interface {
	GetHttpCountLog(ctx context.Context) ([]model.LogCount, error)
}
