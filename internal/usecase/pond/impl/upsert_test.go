package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"testing"
)

func Test_usecase_UpsertPond(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)
	rDbPondMock := new(rDbPond.MockRepository)

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
			name: "success insert new",
			args: args{
				ctx: context.Background(),
				req: model.UpsertPond{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, nil).Once()
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, nil).Once()
				rDbPondMock.On("InsertPond", context.Background(), model.InsertPondReq{PondName: "pond", Description: "desc", FarmID: 1}).Return(nil).Once()
			},
		},
		{
			name: "success update",
			args: args{
				ctx: context.Background(),
				req: model.UpsertPond{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, nil).Once()
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{{PondID: 1}}, nil).Once()
				rDbPondMock.On("UpdatePond", context.Background(), model.UpsertPond{PondID: 1, FarmID: 1, PondName: "pond", Description: "desc"}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbFarm: rDbFarmMock,
				rDbPond: rDbPondMock,
			}
			tt.mock()
			if err := u.UpsertPond(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.UpsertPond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
