package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"errors"
	"testing"
)

func Test_usecase_InsertFarm(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)
	rDbPondMock := new(rDbPond.MockRepository)

	type args struct {
		ctx context.Context
		req model.InsertFarmReq
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
				req: model.InsertFarmReq{
					FarmName:    "farm",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByName", context.Background(), model.GetFarmByNameReq{FarmName: "farm"}).Return([]model.GetFarmByNameRes{}, nil).Once()
				rDbFarmMock.On("InsertFarm", context.Background(), model.InsertFarmReq{FarmName: "farm", Description: "desc"}).Return(nil).Once()
			},
		},
		{
			name: "err GetFarmByName",
			args: args{
				ctx: context.Background(),
				req: model.InsertFarmReq{
					FarmName:    "farm",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByName", context.Background(), model.GetFarmByNameReq{FarmName: "farm"}).Return([]model.GetFarmByNameRes{}, errors.New("err")).Once()
			},
			wantErr: true,
		},
		{
			name: "err duplicate",
			args: args{
				ctx: context.Background(),
				req: model.InsertFarmReq{
					FarmName:    "farm",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByName", context.Background(), model.GetFarmByNameReq{FarmName: "farm"}).Return([]model.GetFarmByNameRes{{FarmID: 1}}, nil).Once()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbFarm: rDbFarmMock,
				rDbPond: rDbPondMock,
			}
			tt.mock()
			if err := u.InsertFarm(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertFarm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
