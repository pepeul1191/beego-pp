package conf

import (
	"github.com/astaxie/beego/context"
)

func BeforeAll(ctx *context.Context) {
	ctx.Output.Header("Server", "beegoServer:1.9.2; Ubuntu")
	ctx.Output.Header("X-Powered-By", "Go")
	ctx.Output.Header("Content-Type", "text/html; charset=UTF-8")
	return
}
