package redisw

import "context"

type Repository interface {
	HIncrBy(ctx context.Context, key string, kv map[string]interface{}) error
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	SAdd(ctx context.Context, key string, members []string) (int64, error)
	SIsMember(ctx context.Context, key string, member string) (bool, error)
	HGet(ctx context.Context, key, field string) (string, error)
	HMSet(ctx context.Context, key string, kv map[string]interface{}) (string, error)
}
