package routers

import (
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"github.com/beego-pp/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Get("/hello", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	beego.Get("/departamento/listar", controllers.DepartamentoListar)
	beego.Router("/departamento", &controllers.DepartamentoController{})
}
