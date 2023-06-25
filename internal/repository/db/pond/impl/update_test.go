package impl

import (
	"aquafarm/common/dbw"
	"aquafarm/common/model"
	"context"
	"errors"
	"testing"
)

func Test_repository_UpdatePond(t *testing.T) {
	dbMock := new(dbw.MockRepository)
	type args struct {
		ctx context.Context
		req model.UpsertPond
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.UpsertPond{
					PondID:      1,
					FarmID:      1,
					Description: "desc",
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryUpdatePond, int64(1), int64(1), "desc").Return(nil, nil).Once()
			},
		},
		{
			name: "err ExecContext",
			args: args{
				ctx: context.Background(),
				req: model.UpsertPond{
					PondID:      1,
					FarmID:      1,
					Description: "desc",
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryUpdatePond, int64(1), int64(1), "desc").Return(nil, errors.New("err")).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			if err := r.UpdatePond(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("repository.UpdatePond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
