package impl

import (
	"aquafarm/common/model"
	rRedisLog "aquafarm/internal/repository/redis/log"
	"context"
	"reflect"
	"testing"
)

func Test_usecase_GetHttpCountLog(t *testing.T) {
	rRedisLogMock := new(rRedisLog.MockRepository)

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
				rRedisLogMock.On("GetHttpCountLog", context.Background()).Return([]model.LogCount{}, nil).Once()
			},
			want: []model.LogCount{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rRedisLog: rRedisLogMock,
			}
			tt.mock()
			got, err := u.GetHttpCountLog(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetHttpCountLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetHttpCountLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
