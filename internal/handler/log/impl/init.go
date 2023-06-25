package impl

import (
	mHttp "aquafarm/common/middleware/http"
	hLog "aquafarm/internal/handler/log"
	uLog "aquafarm/internal/usecase/log"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	uLog  uLog.Usecase
	mHttp mHttp.Middleware
}

func New(httpRouter *httprouter.Router, uLog uLog.Usecase, mHttp mHttp.Middleware) hLog.Handler {
	h := &handler{
		uLog:  uLog,
		mHttp: mHttp,
	}

	// init endpoint
	httpRouter.GET("/log", mHttp.LogWrapper(h.HandlerGetFarmByID))

	return h
}
