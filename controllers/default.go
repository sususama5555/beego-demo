package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Test()  {
	c.Ctx.WriteString("hello")
	//c.Data["result"] = true
	//c.Data["data"] = map[interface{}]int{
	//	"1": 1,
	//	3:3,
	//}
}
