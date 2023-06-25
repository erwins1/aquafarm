package model

type JSONResponse struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}

type LogCount struct {
	Count           int64  `json:"count"`
	UniqueUserAgent int64  `json:"unique_user_agent"`
	URL             string `json:"url"`
	Method          string `json:"method"`
}
