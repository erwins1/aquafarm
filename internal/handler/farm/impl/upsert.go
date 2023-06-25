package impl

import (
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) HandlerUpsertFarm(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp          model.JSONResponse
		err           error
		upsertFarmReq model.UpsertFarm
	)

	// parse request param
	err = json.NewDecoder(r.Body).Decode(&upsertFarmReq)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// validate request
	err = upsertFarmReq.Validate()
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// insert to db
	err = h.uFarm.UpsertFarm(context.Background(), upsertFarmReq)
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
