package impl

import (
	mHttp "aquafarm/common/middleware/http"
	"aquafarm/common/model"
	uFarm "aquafarm/internal/usecase/farm"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerGetFarm(t *testing.T) {
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
				r: &http.Request{
					URL: &url.URL{
						RawQuery: "page=1&per_page=10",
					},
				},
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uFarmMock.On("GetFarm", context.Background(), model.GetFarmReq{Page: 1, PerPage: 10}).Return([]model.GetFarmResult{}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uFarm: uFarmMock,
			}
			tt.mock()
			h.HandlerGetFarm(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}

func Test_handler_HandlerGetFarmByID(t *testing.T) {
	type fields struct {
		uFarm uFarm.Usecase
		mHttp mHttp.Middleware
	}
	type args struct {
		w      http.ResponseWriter
		r      *http.Request
		params httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uFarm: tt.fields.uFarm,
				mHttp: tt.fields.mHttp,
			}
			h.HandlerGetFarmByID(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
