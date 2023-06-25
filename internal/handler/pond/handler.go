package pond

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	HandlerAddPond(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerGetPond(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerGetPondByID(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	HandlerDeletePond(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
