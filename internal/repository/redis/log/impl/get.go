package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"context"
	"encoding/json"
)

func (r *repository) GetHttpCountLog(ctx context.Context) ([]model.LogCount, error) {
	var (
		result = make([]model.LogCount, 0)
		err    error
	)
	logCountDataKey := constant.HttpLogCountRedisKey
	redisData, err := r.redisConnection.HGetAll(ctx, logCountDataKey)

	for _, val := range redisData {
		temp := model.LogCount{}
		if val != "" {
			json.Unmarshal([]byte(val), &temp)
		}
		result = append(result, temp)
	}

	return result, err
}
