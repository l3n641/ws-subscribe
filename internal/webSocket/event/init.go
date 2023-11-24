package event

import (
	"ws/internal/webSocket"
)

func Register() {
	webSocket.WebSocketRouters.AddRouter("setChannel", SetChannel)
}
