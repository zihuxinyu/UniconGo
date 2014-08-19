package Library

import (
	"reflect"
	"github.com/astaxie/beego/orm"
	"encoding/json"
	"fmt"
	"time"
	"strings"
)

type MiniuiGrid struct {
	Total int64 `json:"total"`
	Data  interface{} `json:"data"`
}

//构造新的struct,接收json 绑定

type Data struct {
	List [] map[string]interface{}
}

func save2DB(StructType reflect.Type, SingleItem map[string]interface{}) {
	o := orm.NewOrm()

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
			}else {
				m[f.Name] = SingleItem[f.Name]
			}
		}
	}

	num, err := o.QueryTable(StructType.Name()).Filter("Guid", SingleItem["Guid"]).Update(m)

	fmt.Println("Affected Num: %d, %s", num, err)
}


func Insert2DB(StructType reflect.Type, SingleItem map[string]interface{}) {

	for i := 0; i < StructType.NumField(); i++ {
		f := StructType.Field(i)
		if SingleItem[f.Name] != nil {
			if f.Type == reflect.TypeOf(time.Now()) {
				//对时间格式进行特殊的处理，进行时区转换，miniui过来的json默认为+08:00
				//处理为go转换string为时间需要的标准时间格式
				ss := fmt.Sprintf("%s", SingleItem[f.Name])
				ss = strings.Replace(ss, "T", " ", -1)
				ss = strings.Replace(ss, "+08:00", " +08:00", -1)
				t, _ := time.Parse("2006-01-02 15:04:05 -07:00 ", ss)
				//				fmt.Println("时间字段", t)
				SingleItem[f.Name] = t
			}
		}
		fmt.Println(f.Name, SingleItem[f.Name])



	}

}


func SaveMiniuiData(h interface{}, data string) {

	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)


	//按struct 遍历得到定义，及得到的值
	//	var h Portal_user
	StructType := reflect.TypeOf(h)
	fmt.Println("显示",StructType.Name(),StructType.Elem())


//	v := reflect.New(StructType.Elem())
//	newA:= v.Interface().( StructType.Type())
//	fmt.Println(newA)

	for _, SingleItem := range s.List {
		if state := SingleItem["_state"]; state != nil {
			if state.(string) == "modified" {
			}
		}
		//fmt.Println(k, v["_state"], v)
		save2DB(StructType, SingleItem)
		Insert2DB(StructType, SingleItem)
	}

}

