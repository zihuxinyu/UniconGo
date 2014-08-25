package controllers

import (

	"github.com/astaxie/beego"
	. "UnionGo/models"
	"github.com/astaxie/beego/orm"
	. "UnionGo/Library"

	"fmt"
)

type OptionController struct {
	BaseController

}

func (this *OptionController) Index() {

	ss := GetOptions()
	beego.Debug(ss["key"])
	this.TplNames = "option.html"
	this.Render()
}

func (this *OptionController) Save() {

	data := `{"list":` + this.GetString("data") + `}`
	h := new(Option)

	h.SaveList(data)

	this.Data["json"] = "ok"
	this.ServeJson()
}



func (this *OptionController) Get() {
//
//	pageIndex	0
//	pageSize	10
//	sortField
//	sortOrder
	var pulist []Option
	o := orm.NewOrm()
	pu := new(Option)
	qs := o.QueryTable(pu)
	qs = qs.Limit(20, 0)
	qs.All(&pulist)
	fmt.Println(qs)
	this.Data["json"] = &MiniuiGrid{1000, &pulist}
	this.ServeJson()

}
