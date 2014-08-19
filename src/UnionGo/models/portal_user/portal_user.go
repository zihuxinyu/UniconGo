package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
	"reflect"
	"encoding/json"
	"fmt"
	"strings"
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

	StructType := reflect.TypeOf(h)


	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {
		if state := SingleItem["_state"]; state != nil {
			m := make(orm.Params)
			for i := 0; i < StructType.NumField(); i++ {
				f := StructType.Field(i)
				//fmt.Println(f.Name, f.Type, reflect.TypeOf( v[f.Name]))
				if SingleItem[f.Name] != nil {
					if f.Type == reflect.TypeOf(time.Now()) {
						//对时间格式进行特殊的处理，进行时区转换，miniui过来的json默认为+08:00
						//处理为go转换string为时间需要的标准时间格式
						ss := fmt.Sprintf("%s", SingleItem[f.Name])
						ss = strings.Replace(ss, "T", " ", -1)
						ss = strings.Replace(ss, "+08:00", " +08:00", -1)
						t, _ := time.Parse("2006-01-02 15:04:05 -07:00 ", ss)

						m[f.Name] = t
						//转换正确的时间回填
						SingleItem[f.Name] = t
						fmt.Println(t)
					}else {
						m[f.Name] = SingleItem[f.Name]
					}
				}
			}
			x, _ := json.Marshal(SingleItem)

			var pu Portal_user
			json.Unmarshal(x, &pu)



			switch state.(string){
			case "modified":
				orm.NewOrm().QueryTable(StructType.Name()).Filter("Guid", SingleItem["Guid"]).Update(m)
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

