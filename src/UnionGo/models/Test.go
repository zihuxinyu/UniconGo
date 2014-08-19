package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//标签内容关系表
type TagPost struct {
	Id         int64
	Tagid      int64 `orm:"index"`
	Postid     int64
	Poststatus int8
	Posttime   time.Time
	Username   string `orm:"unique;size(15)"`
	Password   string `orm:"size(32)"`
	Email      string `orm:"size(50)"`
	Lastlogin  time.Time `orm:"auto_now_add;type(datetime)"`
	Logincount int64
	Lastip     string `orm:"size(32)"`
	Authkey    string `orm:"size(10)"`
}

func (m *TagPost) TableName() string {
	return TableName("tag_post")
}
func (m *TagPost) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *TagPost) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *TagPost) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}
func (m *TagPost) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *TagPost) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
