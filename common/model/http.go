package model

type JSONResponse struct {
	Data   interface{} `json:"data"`
	Errors []string    `json:"errors"`
}
