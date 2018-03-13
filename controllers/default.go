package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.mex"
	c.Data["Email"] = "astaxidsfsdfsfsdfe@gmail.com"
	c.TplName = "index.tpl"
}
