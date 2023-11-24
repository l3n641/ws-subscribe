package webSocket

import (
	"encoding/json"
	"fmt"
	"github.com/olahol/melody"
	"sync"
)

var instance *melody.Melody
var once sync.Once

func GetWebsocket() *melody.Melody {
	once.Do(func() {
		instance = melody.New()
		instance.HandleConnect(handleConnect)
		instance.HandleMessage(handleMessage)
		instance.HandleDisconnect(handleDisconnect)
	})
	return instance
}

func handleConnect(s *melody.Session) {
}

func handleMessage(s *melody.Session, request []byte) {
	var response ResponseModel
	data := new(RequestModel)
	err := json.Unmarshal(request, data)
	if err != nil {
		response = ResponseModel{
			Action: "",
			Code:   400,
			Msg:    "解析请求失败",
		}
		s.Write(response.ToByte())
		return
	}

	f, has := WebSocketRouters.GetRouter(data.Action)
	if has != true {
		response = ResponseModel{
			Action: "",
			Code:   400,
			Msg:    "undefined action",
		}
		s.Write(response.ToByte())
		return
	}

	jsonBytes, _ := json.Marshal(data.Data)

	result, err := f(s, jsonBytes)

	if err != nil {
		response = ResponseModel{
			Action: data.Action,
			Code:   500,
			Msg:    err.Error(),
		}
		s.Write(response.ToByte())
	}

	if result != nil {
		response = ResponseModel{
			Action: data.Action,
			Code:   0,
			Msg:    "ok",
			Data:   result,
		}
		s.Write(response.ToByte())
	}

}

func handleDisconnect(s *melody.Session) {
	fmt.Print("断开链接")
}
