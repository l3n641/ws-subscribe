package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"ws/internal/controller/subscribe"
	"ws/internal/webSocket"
	"ws/internal/webSocket/event"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			m := webSocket.GetWebsocket()
			event.Register()

			s.BindHandler("/ws", func(r *ghttp.Request) {
				m.HandleRequest(r.Response.Writer, r.Request)
			})

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(subscribe.Broadcast)
			})

			s.Run()
			return nil
		},
	}
)
