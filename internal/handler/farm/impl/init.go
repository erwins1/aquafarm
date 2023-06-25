package impl

import (
	mHttp "aquafarm/common/middleware/http"
	hFarm "aquafarm/internal/handler/farm"
	uFarm "aquafarm/internal/usecase/farm"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	uFarm uFarm.Usecase
	mHttp mHttp.Middleware
}

func New(httpRouter *httprouter.Router, uFarm uFarm.Usecase, mHttp mHttp.Middleware) hFarm.Handler {
	h := &handler{
		uFarm: uFarm,
		mHttp: mHttp,
	}

	// init endpoint
	httpRouter.GET("/farm", mHttp.LogWrapper(h.HandlerGetFarm))
	httpRouter.GET("/farm/:farm_id", mHttp.LogWrapper(h.HandlerGetFarmByID))
	httpRouter.POST("/farm", mHttp.LogWrapper(h.HandlerAddFarm))
	httpRouter.DELETE("/farm", mHttp.LogWrapper(h.HandlerDeleteFarm))
	httpRouter.PUT("/farm", mHttp.LogWrapper(h.HandlerUpsertFarm))

	return h
}
