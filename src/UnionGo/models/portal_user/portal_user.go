package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
	"reflect"
	"encoding/json"
	"fmt"
	"strconv"
	//. "github.com/mitchellh/mapstructure"
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
	var bb Portal_user
	ModelCache.Set("p", reflect.TypeOf(bb))
}



func Test(h interface{}) {
	mutable := reflect.ValueOf(&h).Elem()
	f := mutable.FieldByName("User_name")
	for i := 0; i < 10; i++ {
		f.SetString("第er次" + strconv.Itoa(i))
		fmt.Println(h)
	}

	kinds := map[string]func() interface{} {

		"p": func() interface{} { return &Portal_user{} },
	}
	fmt.Println(kinds["p"]())

}

var kinds = map[string]func() interface{}{
	"p":func() interface{} {return &Portal_user{}},
}

func ParseForm(form map[string]interface{}, obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)

	objT = objT.Elem()
	objV = objV.Elem()

	for i := 0; i < objT.NumField(); i++ {
		fieldV := objV.Field(i)
		if !fieldV.CanSet() {
			continue
		}

		fieldT := objT.Field(i)
		tags := strings.Split(fieldT.Tag.Get("form"), ",")
		var tag string
		if len(tags) == 0 || len(tags[0]) == 0 {
			tag = fieldT.Name
		} else if tags[0] == "-" {
			continue
		} else {
			tag = tags[0]
		}
		value, ok := form[tag].(string);
		//value := form[tag].(string)
		if len(value) == 0 || !ok {
			continue
		}

		switch fieldT.Type.Kind() {
		case reflect.Bool:
			b, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			fieldV.SetBool(b)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			x, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetInt(x)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			x, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			fieldV.SetUint(x)
		case reflect.Float32, reflect.Float64:
			x, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			fieldV.SetFloat(x)
		case reflect.Interface:
			fieldV.Set(reflect.ValueOf(value))
		case reflect.String:
			fieldV.SetString(value)
		}
	}
	return nil
}

func (h Portal_user) SaveList(data string) {


//	form := map[string]interface{} {
//		"User_name":"魏宝辉",
//		"User_code":"weibh",
//	}

	reflectx := kinds["p"]()
//	hh := Portal_user{}

	//fmt.Println(&hh, reflectx)
	//ParseForm(form, &hh)
	//ParseForm(form, reflectx)

	//fmt.Println(hh, reflectx)




	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)



	//var pu Portal_user

	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range s.List {
		//		fmt.Println(form)
		//		fmt.Println(SingleItem)

		if state := SingleItem["_state"]; state != nil {
			fmt.Println(SingleItem)
			m := make(orm.Params)
			MiniUIDataUpdate(reflectx, SingleItem, m)
			fmt.Println(SingleItem)

			ParseForm(SingleItem, reflectx)
			fmt.Println(reflectx)

			switch state.(string){
			case "modified":
				pk := GetModelPk(reflectx)
				fmt.Println(pk)
				orm.NewOrm().QueryTable(reflect.TypeOf(reflectx).Name()).Filter(pk, SingleItem[pk]).Update(m)
			case "added":
				orm.NewOrm().Insert(reflectx)
				//pu.Insert()
			case "removed":
				orm.NewOrm().Delete(reflectx)
			}
		}
	}
}

//func (h Portal_user) SaveList(data string) {
//
//
//	form := map[string]interface {} {
//		"User_name":"魏宝辉",
//		"User_code":"weibh",
//	}
//
//	reflectx := kinds["p"]()
//	hh := Portal_user{}
//
//	//fmt.Println(&hh, reflectx)
//	ParseForm(form,&hh)
//	ParseForm(form,reflectx)
//
//	fmt.Println(hh,reflectx)
//
//	//fmt.Println(reflect.TypeOf(reflectx).Kind(), reflect.TypeOf(h))
//	orm.NewOrm().Insert(reflectx)
//
//
//	//整理为可识别格式
//	var s Data
//	json.Unmarshal([]byte(data), &s)
//
//
//
//	//var pu Portal_user
//
//	//按struct 遍历得到定义，及得到的值
//	for _, SingleItem := range s.List {
//		//		fmt.Println(form)
//		//		fmt.Println(SingleItem)
//		////		ParseForm(SingleItem,&hh)
//		////		ParseForm(SingleItem,reflectx)
//		//
//		//		fmt.Println(hh,reflectx)
//		if state := SingleItem["_state"]; state != nil {
//			//更新时间格式，并将赋值的字段填充到返回值(orm.Params)中
//			m := make(orm.Params)
//			MiniUIDataUpdate(reflectx, SingleItem, m)
//
//
//			sss := reflectx
//			err := Decode(SingleItem, &sss)
//			if err != nil {
//				panic(err)
//			}
//			fmt.Println("Decode", sss)
//
//			fmt.Println("reflectx", reflectx, reflect.TypeOf(reflectx), reflect.TypeOf(reflectx).Kind())
//
//			switch state.(string){
//			case "modified":
//				pk := GetModelPk(reflectx)
//				orm.NewOrm().QueryTable(reflect.TypeOf(reflectx).Name()).Filter(pk, SingleItem[pk]).Update(m)
//			case "added":
//
//				//pu.Insert()
//			case "removed":
//				//pu.Delete()
//			}
//		}
//	}
//}
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

