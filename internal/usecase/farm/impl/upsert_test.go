package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	"context"
	"testing"
)

func Test_usecase_UpsertFarm(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)

	type args struct {
		ctx context.Context
		req model.UpsertFarm
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		wantErr bool
	}{
		{
			name: "success insert",
			args: args{
				ctx: context.Background(),
				req: model.UpsertFarm{
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
			name: "success update",
			args: args{
				ctx: context.Background(),
				req: model.UpsertFarm{
					FarmName:    "farm",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByName", context.Background(), model.GetFarmByNameReq{FarmName: "farm"}).Return([]model.GetFarmByNameRes{{FarmID: 1}}, nil).Once()
				rDbFarmMock.On("UpdateFarm", context.Background(), model.UpsertFarm{FarmID: 1, FarmName: "farm", Description: "desc"}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbFarm: rDbFarmMock,
			}
			tt.mock()
			if err := u.UpsertFarm(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.UpsertFarm() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
