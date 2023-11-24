package webSocket

import "encoding/json"

type RequestModel struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

type ResponseModel struct {
	Action string      `json:"action"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}

func (r ResponseModel) ToByte() []byte {
	jsonData, _ := json.Marshal(r)
	return jsonData
}
