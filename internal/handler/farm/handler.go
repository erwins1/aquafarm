package farm

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler interface {
	HandlerAddFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
