package impl

import (
	rRedisLog "aquafarm/internal/repository/redis/log"
	uLog "aquafarm/internal/usecase/log"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		rRedisLog rRedisLog.Repository
	}
	tests := []struct {
		name string
		args args
		want uLog.Usecase
	}{
		{
			name: "success",
			args: args{
				rRedisLog: &rRedisLog.MockRepository{},
			},
			want: &usecase{
				rRedisLog: &rRedisLog.MockRepository{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.rRedisLog); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
