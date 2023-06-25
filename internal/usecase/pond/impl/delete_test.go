package impl

import (
	"aquafarm/common/model"
	rDbPond "aquafarm/internal/repository/db/pond"
	"context"
	"testing"
)

func Test_usecase_DeletePondByID(t *testing.T) {
	rDbPondMock := new(rDbPond.MockRepository)

	type args struct {
		ctx context.Context
		req model.DeletePondByID
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
				req: model.DeletePondByID{
					PondID: 1,
				},
			},
			mock: func() {
				rDbPondMock.On("DeletPondByID", context.Background(), model.DeletePondByID{PondID: 1}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &usecase{
				rDbPond: rDbPondMock,
			}
			tt.mock()
			if err := u.DeletePondByID(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("usecase.DeletePondByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
