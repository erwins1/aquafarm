package impl

import (
	"aquafarm/common/dbw"
	"aquafarm/common/model"
	"context"
	"errors"
	"testing"
)

func Test_repository_InsertFarm(t *testing.T) {
	dbMock := new(dbw.MockRepository)

	type args struct {
		ctx  context.Context
		data model.InsertFarmReq
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
				data: model.InsertFarmReq{
					FarmName:    "name",
					Description: "desc",
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryInsertFarm, "name", "desc").Return(nil, nil).Once()
			},
		},
		{
			name: "err ExecContext",
			args: args{
				ctx: context.Background(),
				data: model.InsertFarmReq{
					FarmName:    "name",
					Description: "desc",
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryInsertFarm, "name", "desc").Return(nil, errors.New("err")).Once()
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
			if err := r.InsertFarm(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("repository.InsertFarm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
