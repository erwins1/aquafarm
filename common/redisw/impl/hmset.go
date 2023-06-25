package impl

import "context"

func (r *repository) HMSet(ctx context.Context, key string, kv map[string]interface{}) (string, error) {
	result := r.redisConn.HMSet(ctx, key, kv)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.String(), nil
}
