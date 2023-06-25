package impl

import (
	"aquafarm/common/dbw"
	"aquafarm/common/model"
	"context"
	"errors"
	"reflect"
	"testing"
)

func Test_repository_GetFarmByName(t *testing.T) {
	dbMock := new(dbw.MockRepository)

	type args struct {
		ctx context.Context
		req model.GetFarmByNameReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.GetFarmByNameRes
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByNameReq{
					FarmName: "farm",
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetFarmByNameRes{}, queryGetFarmByName, "farm").Return(nil).Once()
			},
			want: []model.GetFarmByNameRes{},
		},
		{
			name: "err SelectContext",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByNameReq{
					FarmName: "farm",
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetFarmByNameRes{}, queryGetFarmByName, "farm").Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    []model.GetFarmByNameRes{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetFarmByName(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetFarmByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetFarmByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetFarm(t *testing.T) {
	dbMock := new(dbw.MockRepository)

	type args struct {
		ctx    context.Context
		filter model.GetFarmReq
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
				filter: model.GetFarmReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetFarmResult{}, queryGetFarm, int64(10), int64(0)).Return(nil).Once()
			},
			want: []model.GetFarmResult{},
		},
		{
			name: "err SelectContext",
			args: args{
				ctx: context.Background(),
				filter: model.GetFarmReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetFarmResult{}, queryGetFarm, int64(10), int64(0)).Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    []model.GetFarmResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetFarm(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetFarm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetFarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetFarmByID(t *testing.T) {
	dbMock := new(dbw.MockRepository)

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
				dbMock.On("GetContext", context.Background(), &model.GetFarmByIDRes{}, queryGetFarmByID, int64(1)).Return(nil).Once()
			},
			want: model.GetFarmByIDRes{},
		},
		{
			name: "err GetContext",
			args: args{
				ctx: context.Background(),
				req: model.GetFarmByIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				dbMock.On("GetContext", context.Background(), &model.GetFarmByIDRes{}, queryGetFarmByID, int64(1)).Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    model.GetFarmByIDRes{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetFarmByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetFarmByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetFarmByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
