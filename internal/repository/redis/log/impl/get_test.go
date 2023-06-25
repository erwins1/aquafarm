package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"aquafarm/common/redisw"
	"context"
	"reflect"
	"testing"
)

func Test_repository_GetHttpCountLog(t *testing.T) {
	redisMock := new(redisw.MockRepository)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.LogCount
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				redisMock.On("HGetAll", context.Background(), constant.HttpLogCountRedisKey).Return(nil, nil).Once()
			},
			want: make([]model.LogCount, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				redisConnection: redisMock,
			}
			tt.mock()
			got, err := r.GetHttpCountLog(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetHttpCountLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetHttpCountLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
