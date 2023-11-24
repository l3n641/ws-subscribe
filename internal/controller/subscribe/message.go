package subscribe

import (
	"context"
	"encoding/json"
	"github.com/olahol/melody"
	"ws/api/subscribe"
	"ws/internal/webSocket"
)

var Broadcast = cBroadcast{}

type cBroadcast struct {
}

func (c *cBroadcast) All(ctx context.Context, req *subscribe.BroadcastReq) (res *subscribe.BroadcastRes, err error) {
	m := webSocket.GetWebsocket()
	b, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	m.Broadcast(b)
	return &subscribe.BroadcastRes{Msg: "发送成功"}, nil
}

func (c *cBroadcast) Group(ctx context.Context, req *subscribe.BroadcastChannelReq) (res *subscribe.BroadcastRes, err error) {
	m := webSocket.GetWebsocket()
	b, err := json.Marshal(req.Data)
	if err != nil {
		return nil, err
	}
	m.BroadcastFilter(b, func(session *melody.Session) bool {
		channelsInterface, has := session.Get("channels")
		if has == false {
			return false
		}

		if channels, ok := channelsInterface.([]string); ok {

			for _, value1 := range req.Channels {
				for _, value2 := range channels {
					if value1 == value2 {
						return true
					}
				}
			}
			return false
		}
		return false
	})

	return &subscribe.BroadcastRes{Msg: "发送成功"}, nil
}
