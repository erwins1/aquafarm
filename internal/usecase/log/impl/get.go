package impl

import (
	"aquafarm/common/model"
	"context"
)

func (u *usecase) GetHttpCountLog(ctx context.Context) ([]model.LogCount, error) {
	return u.rRedisLog.GetHttpCountLog(ctx)
}
