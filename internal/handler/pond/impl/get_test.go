package impl

import (
	"aquafarm/common/model"
	uPond "aquafarm/internal/usecase/pond"
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerGetPond(t *testing.T) {
	uPondMock := new(uPond.MockUsecase)

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
				uPondMock.On("GetPond", context.Background(), model.GetPondReq{Page: 1, PerPage: 10}).Return([]model.Pond{}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uPond: uPondMock,
			}
			tt.mock()
			h.HandlerGetPond(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
