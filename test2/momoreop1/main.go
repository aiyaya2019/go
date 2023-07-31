package main

import (
	_ "momoreop1/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"momoreop1/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
