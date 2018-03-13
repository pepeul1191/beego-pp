package controllers

import (
	"github.com/astaxie/beego"
	resty "gopkg.in/resty.v1"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["HeadTitle"] = "Home"
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

func (c *MainController) Post() {
	resp, err := resty.R().Get("http://localhost:3000/departamento/listar")
	if err != nil {
		c.Ctx.Output.Status = 500
		c.Ctx.Output.Body([]byte("Error: No se puede conectar contra el servicio"))
	} else {
		c.Ctx.Output.Body([]byte(resp.String()))
	}
}
