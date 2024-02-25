package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mssql/v2"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "one2.3/internal/packed"

	_ "one2.3/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
	"one2.3/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
