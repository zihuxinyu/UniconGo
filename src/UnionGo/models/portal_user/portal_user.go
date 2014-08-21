package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
	"reflect"
	"encoding/json"
	. "github.com/mitchellh/mapstructure"
)

type Portal_user struct{
	Guid                   int `orm:"pk" `
	User_code              string
	User_name              string
	Dpt_name               string
	Msgexpdate             time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Portal_user))

}

var kinds = map[string]func() interface{}{
	"p":func() interface{} {return &Portal_user{}},
}

func Saveit(data string) {
	reflectx := kinds["p"]()


	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)



	//var pu Portal_user

	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {

		if state := SingleItem["_state"]; state != "" {

			m := make(orm.Params)
			//格式化miniui过来的时间
			MiniUIDataUpdate(reflectx, SingleItem, m)

			Decode(SingleItem, reflectx)

			switch state{
			case "modified":
				pk := GetModelPk(reflectx)
				orm.NewOrm().QueryTable(reflect.TypeOf(reflectx).Elem().Name()).Filter(pk, SingleItem[pk]).Update(m)
			case "added":
				orm.NewOrm().Insert(reflectx)
			case "removed":
				orm.NewOrm().Delete(reflectx)
			}
		}
	}
}

func (h Portal_user) SaveList(data string) {
	Saveit(data)
}
