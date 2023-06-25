package impl

import (
	"aquafarm/common/model"
	uLog "aquafarm/internal/usecase/log"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerGetFarmByID(t *testing.T) {
	uLogMock := new(uLog.MockUsecase)

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
						RawQuery: "",
					},
				},
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uLogMock.On("GetHttpCountLog", context.Background()).Return([]model.LogCount{}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uLog: uLogMock,
			}
			tt.mock()
			h.HandlerGetFarmByID(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
