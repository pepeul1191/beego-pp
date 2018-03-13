package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	resty "gopkg.in/resty.v1"
)

type DepartamentoController struct {
	beego.Controller
}

func (c *DepartamentoController) Get() {
	c.Data["HeadTitle"] = "Departamentos"
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["CSSs"] = []string{
		"bower_components/bootstrap/dist/css/bootstrap.min",
		"bower_components/font-awesome/css/font-awesome.min",
	}
	c.Data["JSs"] = []string{
		"bower_components/jquery/dist/jquery.min",
		"bower_components/bootstrap/dist/js/bootstrap.min",
	}
	c.Layout = "layouts/blank.tpl"
	c.TplName = "home/index.tpl"
}

func DepartamentoListar(ctx *context.Context) {
	resp, err := resty.R().Get("http://localhost:3000/departamento/listar")
	if err != nil {
		ctx.Output.Status = 500
		ctx.Output.Body([]byte("Error: No se puede conectar contra el servicio"))
	} else {
		ctx.Output.Body([]byte(resp.String()))
	}
}
