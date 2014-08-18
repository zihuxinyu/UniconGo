package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "quickstart/Library"
	. "quickstart/models/portal_user"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Index() {
	this.TplNames = "index.html"
	this.Render()
}

func (this *MainController) Save() {
	data := `{"list":` + this.GetString("data") + `}`
	h := new(Portal_user)

	h.SaveList(data)
	this.Data["json"] = "ok"
	this.ServeJson()
}

func (this *MainController) Get() {

	var pulist []Portal_user
	o := orm.NewOrm()
	pu := new(Portal_user)

	qs := o.QueryTable(pu)
	qs = qs.Limit(2, 10)
	qs.All(&pulist)

	this.Data["json"] = &MiniuiGrid{1000, &pulist}
	this.ServeJson()

}
