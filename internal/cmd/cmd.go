package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.BindHandler("/ws", func(r *ghttp.Request) {
				ws, err := r.WebSocket()
				if err != nil {
					glog.Error(ctx, err)
					r.Exit()
				}
				for {
					msgType, msg, err := ws.ReadMessage()
					if err != nil {
						return
					}
					if err = ws.WriteMessage(msgType, msg); err != nil {
						return
					}
				}
			})
			s.SetServerRoot(gfile.MainPkgPath())
			s.SetPort(8199)
			s.Run()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind()
			})
			s.Run()
			return nil
		},
	}
)
