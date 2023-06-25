package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"errors"
	"reflect"
	"testing"
)

func Test_usecase_GetFarm(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)
	rDbPondMock := new(rDbPond.MockRepository)

	type args struct {
		ctx context.Context
		req model.GetFarmReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.GetFarmResult
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarm", context.Background(), model.GetFarmReq{Page: 1, PerPage: 10}).Return([]model.GetFarmResult{}, nil).Once()
			},
			want: []model.GetFarmResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbFarm: rDbFarmMock,
				rDbPond: rDbPondMock,
			}
			tt.mock()
			got, err := u.GetFarm(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetFarm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_GetFarmByID(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)
	rDbPondMock := new(rDbPond.MockRepository)
	type args struct {
		ctx context.Context
		req model.GetFarmByIDReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    model.GetFarmByIDRes
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{FarmID: 1}, nil).Once()
				rDbPondMock.On("GetPondByFarmID", context.Background(), model.GetPondByFarmIDReq{FarmID: 1}).Return([]model.Pond{}, nil).Once()
			},
			want: model.GetFarmByIDRes{
				FarmID: 1,
				Ponds:  []model.Pond{},
			},
		},
		{
			name: "err GetFarmByID",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{}, errors.New("err")).Once()
			},
			wantErr: true,
		},
		{
			name: "err not found",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				rDbFarmMock.On("GetFarmByID", context.Background(), model.GetFarmByIDReq{FarmID: 1}).Return(model.GetFarmByIDRes{}, nil).Once()
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
			got, err := u.GetFarmByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetFarmByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetFarmByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
