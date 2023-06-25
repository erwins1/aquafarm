package impl

import "context"

func (r *repository) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	redisResult := r.redisConn.HGetAll(ctx, key)
	if redisResult.Err() != nil {
		return nil, redisResult.Err()
	}

	return redisResult.Val(), nil
}

func (r *repository) HGet(ctx context.Context, key, field string) (string, error) {
	result := r.redisConn.HGet(ctx, key, field)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Result()
}
