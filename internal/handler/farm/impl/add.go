package impl

import (
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) HandlerAddFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp          model.JSONResponse
		err           error
		insertFarmReq model.InsertFarmReq
	)

	// parse request param
	err = json.NewDecoder(r.Body).Decode(&insertFarmReq)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// validate request
	err = insertFarmReq.Validate()
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// insert to db
	err = h.uFarm.InsertFarm(context.Background(), insertFarmReq)
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	resp.Data = "success"
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}
