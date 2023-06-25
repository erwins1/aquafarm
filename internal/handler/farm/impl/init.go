package impl

import (
	hFarm "aquafarm/internal/handler/farm"
	uFarm "aquafarm/internal/usecase/farm"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	uFarm uFarm.Usecase
}

func New(httpRouter *httprouter.Router, uFarm uFarm.Usecase) hFarm.Handler {
	h := &handler{
		uFarm: uFarm,
	}

	// init endpoint
	httpRouter.GET("/farm", h.HandlerGetFarm)
	httpRouter.GET("/farm/:farm_id", h.HandlerGetFarmByID)
	httpRouter.POST("/farm", h.HandlerAddFarm)
	httpRouter.DELETE("/farm", h.HandlerDeleteFarm)
	httpRouter.PUT("/farm", h.HandlerUpsertFarm)

	return h
}
