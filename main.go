package main

import (
	"github.com/go-mogu/mogu-gateway/internal/cmd"
	_ "github.com/go-mogu/mogu-gateway/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
