package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["title"] = "Home"
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
