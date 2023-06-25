package farm

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	HandlerAddFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerGetFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerGetFarmByID(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerDeleteFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
