package impl

import (
	mHttp "aquafarm/common/middleware/http"
	hPond "aquafarm/internal/handler/pond"
	uPond "aquafarm/internal/usecase/pond"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	uPond uPond.Usecase
	mHttp mHttp.Middleware
}

func New(httpRouter *httprouter.Router, uPond uPond.Usecase, mHttp mHttp.Middleware) hPond.Handler {
	h := &handler{
		uPond: uPond,
	}

	// init endpoint
	httpRouter.GET("/pond", mHttp.LogWrapper(h.HandlerGetPond))
	httpRouter.GET("/pond/:pond_id", mHttp.LogWrapper(h.HandlerGetPondByID))
	httpRouter.POST("/pond", mHttp.LogWrapper(h.HandlerAddPond))
	httpRouter.DELETE("/pond", mHttp.LogWrapper(h.HandlerDeletePond))
	httpRouter.PUT("/pond", mHttp.LogWrapper(h.HandlerUpsertPond))

	return h
}
