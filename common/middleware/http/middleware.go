package http

import "github.com/julienschmidt/httprouter"

type Middleware interface {
	LogWrapper(next httprouter.Handle) httprouter.Handle
}
