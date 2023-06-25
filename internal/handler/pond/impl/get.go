package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) HandlerGetPond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp       model.JSONResponse
		err        error
		getPondReq model.GetPondReq
		getPondRes []model.Pond
	)

	// parse request param
	err = getPondReq.ParseAndValidate(r)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// insert to db
	getPondRes, err = h.uPond.GetPond(context.Background(), getPondReq)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	resp.Data = getPondRes
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}

func (h *handler) HandlerGetPondByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp           model.JSONResponse
		err            error
		getPondByIDReq model.GetPondByID
		getPondByIDRes model.Pond
	)

	// parse request param
	getPondByIDReq.PondID, err = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// get data
	getPondByIDRes, err = h.uPond.GetPondByID(context.Background(), getPondByIDReq)
	if err != nil {
		if err.Error() == constant.ErrNotFound { // check if the id not exist
			w.WriteHeader(http.StatusNotFound)
			resp.Errors = append(resp.Errors, err.Error())
			result, _ := json.Marshal(resp)
			fmt.Fprint(w, string(result))
			return
		}
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// returning result
	resp.Data = getPondByIDRes
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}
