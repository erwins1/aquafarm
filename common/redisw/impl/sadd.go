package impl

import (
	"context"
)

func (r *repository) SAdd(ctx context.Context, key string, members []string) (int64, error) {

	membersIntf := SliceOfStringsToSliceOfInterfaces(members)
	result, err := r.redisConn.SAdd(ctx, key, membersIntf...).Result()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *repository) SIsMember(ctx context.Context, key string, member string) (bool, error) {

	memberIntf := ConvertStringToInterface(member)
	result, err := r.redisConn.SIsMember(ctx, key, memberIntf).Result()
	if err != nil {
		return result, err
	}

	return result, nil
}
