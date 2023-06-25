package impl

import "context"

func (r *repository) HIncrBy(ctx context.Context, key string, kv map[string]interface{}) error {
	var (
		err error
	)

	for field, value := range kv {
		cmd := r.redisConn.HIncrBy(ctx, key, field, value.(int64))
		if cmd.Err() != nil {
			return cmd.Err()
		}
	}

	return err
}
