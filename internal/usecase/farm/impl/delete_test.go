package impl

import (
	"aquafarm/common/model"
	rDbFarm "aquafarm/internal/repository/db/farm"
	"context"
	"testing"
)

func Test_usecase_DeleteFarmByID(t *testing.T) {
	rDbFarmMock := new(rDbFarm.MockRepository)

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
				rDbFarmMock.On("DeleteFarmByID", context.Background(), model.DeleteFarmByIDReq{FarmID: 1}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbFarm: rDbFarmMock,
			}
			tt.mock()
			if err := u.DeleteFarmByID(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.DeleteFarmByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
