package main

import (
	"github.com/astaxie/beego"
	"github.com/beego-pp/conf"
	_ "github.com/beego-pp/routers"
)

func main() {
	// set contantes
	conf.SetConstants()
	// cargar conf/helpers.go
	beego.AddFuncMap("LoadCSSs", conf.LoadCSSs)
	beego.AddFuncMap("LoadJSs", conf.LoadJSs)
	// cargar conf/middelware.go
	beego.InsertFilter("/*", beego.BeforeRouter, conf.BeforeAll)
	// run
	beego.Run()
}
