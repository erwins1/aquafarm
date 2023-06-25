package impl

import (
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *handler) HandlerGetFarmByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var (
		resp           model.JSONResponse
		err            error
		getLogCountRes []model.LogCount
	)

	getLogCountRes, err = h.uLog.GetHttpCountLog(context.Background())
	if err != nil {
		resp.Errors = append(resp.Errors, err.Error())
		result, _ := json.Marshal(resp)
		fmt.Fprint(w, string(result))
		return
	}

	resp.Data = getLogCountRes
	result, _ := json.Marshal(resp)
	fmt.Fprint(w, string(result))
}
