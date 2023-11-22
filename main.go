package main

import (
	_ "ws/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"ws/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
