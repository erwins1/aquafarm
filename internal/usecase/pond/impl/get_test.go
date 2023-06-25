package impl

import (
	"aquafarm/common/model"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"reflect"
	"testing"
)

func Test_usecase_GetPond(t *testing.T) {
	rDbPondMock := new(rDbPond.MockRepository)

	type args struct {
		ctx context.Context
		req model.GetPondReq
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
				req: model.GetPondReq{
					Page:    1,
					PerPage: 10,
				},
			},
			mock: func() {
				rDbPondMock.On("GetPond", context.Background(), model.GetPondReq{Page: 1, PerPage: 10}).Return([]model.Pond{}, nil).Once()
			},
			want: []model.Pond{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbPond: rDbPondMock,
			}
			tt.mock()
			got, err := u.GetPond(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetPond() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetPond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_GetPondByID(t *testing.T) {
	rDbPondMock := new(rDbPond.MockRepository)

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
				rDbPondMock.On("GetPondByID", context.Background(), model.GetPondByID{PondID: 1}).Return(model.Pond{PondID: 1}, nil).Once()
			},
			want: model.Pond{
				PondID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbPond: rDbPondMock,
			}
			tt.mock()
			got, err := u.GetPondByID(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("usecase.GetPondByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("usecase.GetPondByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
