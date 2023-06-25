package impl

import (
	"aquafarm/common/redisw"
	rRedisLog "aquafarm/internal/repository/redis/log"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		redisConnection redisw.Repository
	}
	tests := []struct {
		name string
		args args
		want rRedisLog.Repository
	}{
		{
			name: "success",
			args: args{
				&redisw.MockRepository{},
			},
			want: &repository{
				&redisw.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.redisConnection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
