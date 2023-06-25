package impl

import (
	"aquafarm/common/constant"
	"aquafarm/common/model"
	"context"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

func (m *middleware) LogWrapper(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		go m.insertHttpCountLog(r.UserAgent(), r.URL.Path, r.Method) // we will do it at background and ignore any error because we dont want log blocking the whole process
		next(w, r, p)
	}
}

func (m *middleware) insertHttpCountLog(userAgent, urlPath, method string) {

	var (
		data            model.LogCount
		uniqueUserCount int64
		hmSetReq        = make(map[string]interface{}, 0)
	)
	// check user agent exist or not
	sMembersRedisKey := fmt.Sprintf(constant.UniqueUserAgentRedisKey, urlPath, method)
	isExist, err := m.redisConnection.SIsMember(context.Background(), sMembersRedisKey, userAgent)
	if err != nil {
		return
	}
	// check user agent is exist or not
	if !isExist {
		uniqueUserCount += 1
		m.redisConnection.SAdd(context.Background(), sMembersRedisKey, []string{userAgent})
	}
	// get existing data
	logCountDataKey := constant.HttpLogCountRedisKey
	result, err := m.redisConnection.HGet(context.Background(), logCountDataKey, fmt.Sprintf("%s-%s", urlPath, method))
	if err != nil && err != redis.Nil {
		return
	}

	if result != "" {
		json.Unmarshal([]byte(result), &data)
	}
	// add count
	data.Count += 1
	data.UniqueUserAgent += uniqueUserCount
	data.URL = urlPath
	data.Method = method

	insertRedisData, err := json.Marshal(data)
	if err != nil {
		return
	}
	// save new data
	hmSetReq[fmt.Sprintf("%s-%s", urlPath, method)] = insertRedisData
	m.redisConnection.HMSet(context.Background(), logCountDataKey, hmSetReq)
}
