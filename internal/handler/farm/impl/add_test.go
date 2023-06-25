package impl

import (
	"aquafarm/common/model"
	uFarm "aquafarm/internal/usecase/farm"
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerAddFarm(t *testing.T) {
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
					req, _ := http.NewRequest("POST", "/farm",
						bytes.NewBufferString(
							`{"farm_name": "farm","description": "desc"}`,
						),
					)
					return req
				}(),
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uFarmMock.On("InsertFarm", context.Background(), model.InsertFarmReq{FarmName: "farm", Description: "desc"}).Return(nil).Once()
			},
		},
		{
			name: "err InsertFarm",
			args: args{
				r: func() *http.Request {
					req, _ := http.NewRequest("POST", "/farm",
						bytes.NewBufferString(
							`{"farm_name": "farm","description": "desc"}`,
						),
					)
					return req
				}(),
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uFarmMock.On("InsertFarm", context.Background(), model.InsertFarmReq{FarmName: "farm", Description: "desc"}).Return(errors.New("err")).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uFarm: uFarmMock,
			}
			tt.mock()
			h.HandlerAddFarm(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
