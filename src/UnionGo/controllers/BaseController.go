package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type BaseController struct {
	beego.Controller
}

func (c BaseController) GetUserId() string {
	if userID := c.GetSession("UserID"); userID != nil {
		return userID.(string)
	}
	return ""
}

// 是否已登录
func (c BaseController) HasLogined() bool {
	return c.GetUserId() != ""
}



func (c BaseController) ClearSession() {
	c.DelSession("UserID")
}

// 修改session
func (c BaseController) UpdateSession(key, value string) {
	c.SetSession(key, value)
}

// 返回json
func (c BaseController) Json(i interface{}) string {
	// b, _ := json.MarshalIndent(i, "", " ")
	b, _ := json.Marshal(i)
	return string(b)
}


