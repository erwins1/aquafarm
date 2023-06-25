package log

import (
	"aquafarm/common/model"
	"context"
)

type Usecase interface {
	GetHttpCountLog(ctx context.Context) ([]model.LogCount, error)
}
