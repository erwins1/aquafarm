package impl

import (
	"aquafarm/common/dbw"
	"aquafarm/common/model"
	"context"
	"errors"
	"reflect"
	"testing"
)

func Test_repository_GetPondByFarmID(t *testing.T) {
	dbMock := new(dbw.MockRepository)

	type args struct {
		ctx context.Context
		req model.GetPondByFarmIDReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.Pond
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByFarmIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.Pond{}, queryGetPondByFarmID, int64(1)).Return(nil).Once()
			},
			want: []model.Pond{},
		},
		{
			name: "err SelectContext",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByFarmIDReq{
					FarmID: 1,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.Pond{}, queryGetPondByFarmID, int64(1)).Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    []model.Pond{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetPondByFarmID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetPondByFarmID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetPondByFarmID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetPondByName(t *testing.T) {
	dbMock := new(dbw.MockRepository)
	type args struct {
		ctx context.Context
		req model.GetPondByNameReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.GetPondByNameRes
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByNameReq{
					PondName: "name",
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetPondByNameRes{}, queryGetPondByName, "name").Return(nil).Once()
			},
			want: []model.GetPondByNameRes{},
		},
		{
			name: "err SelectContext",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByNameReq{
					PondName: "name",
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.GetPondByNameRes{}, queryGetPondByName, "name").Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    []model.GetPondByNameRes{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetPondByName(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetPondByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetPondByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetPond(t *testing.T) {
	dbMock := new(dbw.MockRepository)
	type args struct {
		ctx    context.Context
		filter model.GetPondReq
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    []model.Pond
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				filter: model.GetPondReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.Pond{}, queryGetPond, int64(10), int64(0)).Return(nil).Once()
			},
			want: []model.Pond{},
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				filter: model.GetPondReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				dbMock.On("SelectContext", context.Background(), &[]model.Pond{}, queryGetPond, int64(10), int64(0)).Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    []model.Pond{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetPond(tt.args.ctx, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetPond() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetPond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repository_GetPondByID(t *testing.T) {
	dbMock := new(dbw.MockRepository)
	type args struct {
		ctx context.Context
		req model.GetPondByID
	}
	tests := []struct {
		name    string
		mock    func()
		args    args
		want    model.Pond
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByID{
					PondID: 1,
				},
			},
			mock: func() {
				dbMock.On("GetContext", context.Background(), &model.Pond{}, queryGetPondByID, int64(1)).Return(nil).Once()
			},
			want: model.Pond{},
		},
		{
			name: "err GetContext",
			args: args{
				ctx: context.Background(),
				req: model.GetPondByID{
					PondID: 1,
				},
			},
			mock: func() {
				dbMock.On("GetContext", context.Background(), &model.Pond{}, queryGetPondByID, int64(1)).Return(errors.New("err")).Once()
			},
			wantErr: true,
			want:    model.Pond{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db: dbMock,
			}
			tt.mock()
			got, err := r.GetPondByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("repository.GetPondByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repository.GetPondByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
