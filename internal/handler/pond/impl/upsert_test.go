package impl

import (
	"aquafarm/common/model"
	uPond "aquafarm/internal/usecase/pond"
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func Test_handler_HandlerUpsertPond(t *testing.T) {
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
				r: func() *http.Request {
					req, _ := http.NewRequest("PUT", "/farm",
						bytes.NewBufferString(
							`{"pond_name": "pond","description": "desc", "farm_id": 1}`,
						),
					)
					return req
				}(),
				w: httptest.NewRecorder(),
			},
			mock: func() {
				uPondMock.On("UpsertPond", context.Background(), model.UpsertPond{FarmID: 1, PondName: "pond", Description: "desc"}).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				uPond: uPondMock,
			}
			tt.mock()
			h.HandlerUpsertPond(tt.args.w, tt.args.r, tt.args.params)
		})
	}
}
