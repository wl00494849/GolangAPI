package model

import "encoding/json"

type ResultModel struct {
	Code     int
	Body     interface{}
	IsSucess bool
}

func (res *ResultModel) ResponseResult() []byte {
	result, _ := json.Marshal(res)
	return result
}
