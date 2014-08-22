package Library

import (
	"strings"
	"reflect"
	"time"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"errors"
)

//找出beegoModel的主键
func GetModelPk(obj interface{}) (pkFiledName string) {
	s := reflect.TypeOf(obj).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		pkFiled := s.Field(i)
		tags := strings.Split(pkFiled.Tag.Get("orm"), ",")
		for _, v := range tags {
			if strings.ToLower(v) == "pk" {
				pkFiledName = pkFiled.Name
				break
			}

		}
		//得到值就退出循环
		if len(pkFiledName) > 0 {break}
	}
	return pkFiledName
}

//格式化miniui过来的时间,得到修改过的字段放到m中
func MiniUIDataUpdate(obj interface{}, SingleItem map[string]interface{}, m orm.Params) error {
	if (!isPtr(obj)) {
		return errors.New(fmt.Sprintf("只支持指针类型，不支持`%T`", obj))
	}
	StructType := reflect.TypeOf(obj).Elem() //通过反射获取type定义

	for i := 0; i < StructType.NumField(); i++ {
		f := StructType.Field(i)
		//fmt.Println(f.Name, f.Type, reflect.TypeOf( v[f.Name]))


		//此处日后可以做根据字段名注入逻辑控制后的字段，比如创建人等信息
		if SingleItem[f.Name] != nil {
			//fmt.Println("格式整理",f.Name,reflect.TypeOf( SingleItem[f.Name]))
			if f.Type == reflect.TypeOf(time.Now()) {
				//对时间格式进行特殊的处理，进行时区转换，miniui过来的时间加+08:00
				//处理为go转换string为时间需要的标准时间格式
				ss := fmt.Sprintf("%s", SingleItem[f.Name])
				ss = strings.Replace(ss, "T", " ", -1)
				ss = ss+" +08:00"
				//fmt.Println("时间格式整理"+ss)
				//ss = strings.Replace(ss, "+08:00", " +08:00", -1)

				t, _ := time.Parse("2006-01-02 15:04:05 -07:00 ", ss)

				m[f.Name] = t
				//转换正确的时间回填
				//SingleItem[f.Name] = t.Format("2006-01-02 15:04:05")
				SingleItem[f.Name] = t
			}else {

				m[f.Name] = SingleItem[f.Name]
			}
			//fmt.Println("格式整理后",f.Name,reflect.TypeOf( SingleItem[f.Name]))

		}
	}

	//先将map 对应为json
	x, _ := json.Marshal(SingleItem)
	//fmt.Println(reflect.TypeOf(x), string(x))
	//再将json对应为struct
	json.Unmarshal(x, obj)
	
	return nil
}
///将miniui过来的数据保存，根据ModelName通过ModelCache模块获得model实例
func SaveMiniUIData(ModelName string ,data string) {
	//根据名称获取Model
	reflecty,_:=ModelCache.Get(ModelName)
	//得到实例
	reflectx :=  reflecty()

	//整理为可识别格式
	var dataList DataList
	json.Unmarshal([]byte(data), &dataList)


	//按struct 遍历得到定义，及得到的值
	for _, SingleItem := range dataList.List {

		if state := SingleItem["_state"]; state != "" {

			m := make(orm.Params)
			//格式化miniui过来的时间,得到修改过的字段放到m中
			MiniUIDataUpdate(reflectx, SingleItem, m)

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
