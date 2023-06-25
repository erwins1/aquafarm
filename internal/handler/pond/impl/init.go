package impl

import (
	hPond "aquafarm/internal/handler/pond"
	uPond "aquafarm/internal/usecase/pond"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	uPond uPond.Usecase
}

func New(httpRouter *httprouter.Router, uPond uPond.Usecase) hPond.Handler {
	h := &handler{
		uPond: uPond,
	}

	// init endpoint
	httpRouter.GET("/pond", h.HandlerGetPond)
	httpRouter.GET("/pond/:pond_id", h.HandlerGetPondByID)
	httpRouter.POST("/pond", h.HandlerAddPond)
	httpRouter.DELETE("/pond", h.HandlerDeletePond)
	httpRouter.PUT("/pond", h.HandlerUpsertPond)

	return h
}
