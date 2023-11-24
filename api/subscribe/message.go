package subscribe

import "github.com/gogf/gf/v2/frame/g"

type BroadcastReq struct {
	g.Meta `path:"/broadcast" method:"post" summary:"广播通知所有人"`
	Data   *g.Var `v:"required#数据必须填写" dc:"消息"`
}

type BroadcastChannelReq struct {
	g.Meta   `path:"/broadcast-channel" method:"post" summary:"广播通知指定通道的用户"`
	Channels []string    `v:"required#频道必须选择" dc:"频道"`
	Data     interface{} `v:"required#数据必须填写" dc:"消息"`
}

type BroadcastRes struct {
	Msg string `json:"msg"`
}
