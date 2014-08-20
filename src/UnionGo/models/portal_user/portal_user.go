package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
	"reflect"
	"encoding/json"

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

func (h Portal_user) SaveList(data string) {
	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)
	pk:=GetModelPk(h)
	StructType := reflect.TypeOf(h)


	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {
		if state := SingleItem["_state"]; state != nil {
			//更新时间格式，并将赋值的字段填充到返回值(orm.Params)中
			m :=MiniUIDataUpdate(StructType, SingleItem)

			var pu Portal_user
			x, _ := json.Marshal(SingleItem)
			json.Unmarshal(x, &h)


			switch state.(string){
			case "modified":
				orm.NewOrm().QueryTable(StructType.Name()).Filter(pk, SingleItem[pk]).Update(m)
			case "added":
				pu.Insert()
			case "removed":
				pu.Delete()
			}
		}
	}
	//SaveMiniuiData(h,data)
}
func (m *Portal_user) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *Portal_user) Delete() error {

	intType := reflect.TypeOf(m).Elem()
	intPtr2 := reflect.New(intType)
	// Just to prove it
	b := intPtr.Elem().Interface().(int)

	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *Portal_user) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

