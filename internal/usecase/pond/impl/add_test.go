package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"errors"
	"testing"
)

func Test_usecase_InsertPond(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)
	rDbPondMock := new(rDbPond.MockRepository)
	type args struct {
		ctx context.Context
		req model.InsertPondReq
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
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, nil).Once()
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, nil).Once()
				rDbPondMock.On("InsertPond", context.Background(), model.InsertPondReq{PondName: "pond", Description: "desc", FarmID: 1}).Return(nil).Once()
			},
		},
		{
			name: "err GetPondByName",
			args: args{
				ctx: context.Background(),
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, errors.New("err")).Once()
			},
			wantErr: true,
		},
		{
			name: "err duplicate",
			args: args{
				ctx: context.Background(),
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{{PondID: 1}}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "err GetFarmByID",
			args: args{
				ctx: context.Background(),
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, nil).Once()
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, errors.New("err")).Once()
			},
			wantErr: true,
		},
		{
			name: "err invalid farm id",
			args: args{
				ctx: context.Background(),
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, nil).Once()
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 0}, nil).Once()
			},
			wantErr: true,
		},
		{
			name: "err InsertPond",
			args: args{
				ctx: context.Background(),
				req: model.InsertPondReq{
					FarmID:      1,
					PondName:    "pond",
					Description: "desc",
				},
			},
			mock: func() {
				rDbPondMock.On("GetPondByName", context.Background(), model.GetPondByNameReq{PondName: "pond"}).Return([]model.GetPondByNameRes{}, nil).Once()
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, nil).Once()
				rDbPondMock.On("InsertPond", context.Background(), model.InsertPondReq{PondName: "pond", Description: "desc", FarmID: 1}).Return(errors.New("err")).Once()
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
			if err := u.InsertPond(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.InsertPond() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
