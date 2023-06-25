package impl

import (
	"aquafarm/common/model"
	uFarm "aquafarm/internal/usecase/farm"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerDeleteFarm(t *testing.T) {
	uFarmMock := new(uFarm.MockUsecase)

	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name string
		mock func()
		args args
	}{
		{
			name: "success",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest("DELETE", "/farm",
						bytes.NewBufferString(
							`{"farm_id": 1}`,
						),
					)
					return req
				}(),
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uFarmMock.On("DeleteFarmByID", context.Background(), model.DeleteFarmByIDReq{FarmID: 1}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uFarm: uFarmMock,
			}
			tt.mock()
			h.HandlerDeleteFarm(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
