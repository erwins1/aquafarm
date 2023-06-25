package impl

import (
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) HandlerDeletePond(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp          model.JSONResponse
		err           error
		deletePondReq model.DeletePondByID
	)

	// parse request param
	err = json.NewDecoder(r.Body).Decode(&deletePondReq)
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// validate request
	err = deletePondReq.Validate()
	if err != nil { // check error
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	// insert to db
	err = h.uPond.DeletePondByID(context.Background(), deletePondReq)
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
