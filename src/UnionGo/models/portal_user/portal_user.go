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
}

func (h Portal_user) SaveList(data string) {

	cc, _ := ModelCache.Get("p")
	intPtr := reflect.New(reflect.TypeOf(cc))

	intV := reflect.ValueOf(&intPtr).Elem()
	fv := intV.FieldByName("User_name")
	fv.SetString("好了吗")

	mutable := reflect.ValueOf(&h).Elem()
	f := mutable.FieldByName("User_name")
	for  i:=0;i<10;i++{
		f.SetString("第一次"+strconv.Itoa(i))
		fmt.Println(h)
	}
	Test(h)


//
//
//	intPtr := reflect.New(reflect.TypeOf(bb))
//	var sss intPtr
//	sss.User_name="dddd"
//	intV:=reflect.ValueOf(&intPtr).Elem()
//	fv:=intV.FieldByName("User_name")
//	fv.SetString("好了吗")
//
//
//
//	cc, _ := ModelCache.Get("p")
//
//
//	xyz := reflect.New(cc).Elem()
//	fmt.Println("real", reflect.TypeOf(bb) )
//	fmt.Println("ModelCache", intPtr.Elem().Interface())
//
//	fmt.Println("xyz反射后得到", xyz)
//	mutable2 := reflect.ValueOf(&xyz).Elem()
//	f2 := mutable2.FieldByName("User_name")
//
//	fmt.Println(cc, "第二次", mutable2.Kind(), f2.IsValid(), f2.CanSet(), f2.Kind())
//	f2.SetString("dddd")
//
//	fmt.Println("第二次", cc)
//
//
//	fmt.Println(ModelCache.Get("p"))

	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)
	pk := GetModelPk(h)
	StructType := reflect.TypeOf(h)


	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {
		if state := SingleItem["_state"]; state != nil {
			//更新时间格式，并将赋值的字段填充到返回值(orm.Params)中
			m := make(orm.Params)
			MiniUIDataUpdate(StructType, SingleItem, m)

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

