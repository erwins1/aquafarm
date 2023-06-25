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

func (h *handler) HandlerGetFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp       model.JSONResponse
		err        error
		getFarmReq model.GetFarmReq
		getFarmRes []model.GetFarmResult
	)

	// parse request param
	err = getFarmReq.ParseAndValidate(r)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// insert to db
	getFarmRes, err = h.uFarm.GetFarm(context.Background(), getFarmReq)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	resp.Data = getFarmRes
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}

func (h *handler) HandlerGetFarmByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp       model.JSONResponse
		err        error
		getFarmReq model.GetFarmByIDReq
		getFarmRes model.GetFarmByIDRes
	)

	// parse request param
	getFarmReq.FarmID, err = strconv.ParseInt(path.Base(r.URL.Path), 10, 64)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	getFarmRes, err = h.uFarm.GetFarmByID(context.Background(), getFarmReq)
	if err != nil {
		if err.Error() == constant.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			resp.Errors = append(resp.Errors, err.Error())
			result, _ := json.Marshal(resp)
			fmt.Fprint(w, string(result))
			return
		}
	}

	resp.Data = getFarmRes
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}
