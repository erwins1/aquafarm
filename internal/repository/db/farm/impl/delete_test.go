package impl

import (
	"aquafarm/common/dbw"
	"aquafarm/common/model"
	"context"
	"errors"
	"testing"
)

func Test_repository_DeleteFarmByID(t *testing.T) {

	dbMock := new(dbw.MockRepository)

	type args struct {
		ctx context.Context
		req model.DeleteFarmByIDReq
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
				req: model.DeleteFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryDeleteFarm, int64(1)).Return(nil, nil).Once()
			},
		},
		{
			name: "err ExecContext",
			args: args{
				ctx: context.Background(),
				req: model.DeleteFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				dbMock.On("ExecContext", context.Background(), queryDeleteFarm, int64(1)).Return(nil, errors.New("err")).Once()
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
			if err := r.DeleteFarmByID(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("repository.DeleteFarmByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
