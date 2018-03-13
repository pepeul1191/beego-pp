package main

import (
	"github.com/astaxie/beego"
	"github.com/beego-pp/conf"
	_ "github.com/beego-pp/routers"
)

func main() {
	conf.SetConstants()
	beego.AddFuncMap("LoadCSSs", conf.LoadCSSs)
	beego.AddFuncMap("LoadJSs", conf.LoadJSs)
	beego.Run()
}
