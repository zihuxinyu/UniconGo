package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
	"reflect"
	"encoding/json"
	"fmt"
	"strconv"
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
	var bb Portal_user
	ModelCache.Set("p", reflect.TypeOf(bb))
}



func Test(h interface {}){
	mutable := reflect.ValueOf(&h).Elem()
	f := mutable.FieldByName("User_name")
	for  i:=0;i<10;i++{
		f.SetString("第er次"+strconv.Itoa(i))
		fmt.Println(h)
	}

	kinds := map[string]func() interface {} {

		"p": func() interface {} { return &Portal_user{} },
	}
	fmt.Println(kinds["p"]())

}
var kinds=map[string]func() interface {}{
	"p":func() interface {}{return Portal_user{}},
}


func (h Portal_user) SaveList(data string) {

	reflectx:=kinds["p"]()


	fmt.Println(reflect.TypeOf(reflectx),reflect.TypeOf(h))

	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)

	StructType := reflect.TypeOf(h)


	//var pu Portal_user

	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {
		if state := SingleItem["_state"]; state != nil {
			//更新时间格式，并将赋值的字段填充到返回值(orm.Params)中
			m := make(orm.Params)
			MiniUIDataUpdate(StructType, SingleItem, m)


			x, _ := json.Marshal(SingleItem)
			json.Unmarshal(x, &reflectx)
			fmt.Println("reflectx",reflectx,reflect.TypeOf(reflectx))
			switch state.(string){
			case "modified":
				pk := GetModelPk(reflectx)
				fmt.Println(pk)
				orm.NewOrm().QueryTable(StructType.Name()).Filter(pk, SingleItem[pk]).Update(m)
			case "added":
				//pu.Insert()
			case "removed":
				//pu.Delete()
			}
		}
	}
}
func (m *Portal_user) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
func (m *Portal_user) Delete() error {
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

