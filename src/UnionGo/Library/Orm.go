package Library

import (
	"strings"
	"reflect"
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
)


//找出beegoModel的主键
func GetModelPk(obj interface{}) (pkFiledName string) {
	s := reflect.TypeOf(obj).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		pkFiled := s.Field(i)
		tags := strings.Split(pkFiled.Tag.Get("orm"), ",")
		for _, v := range tags {
			if strings.ToLower(v) == "pk" {
				pkFiledName= pkFiled.Name
				break
			}

		}
		//得到值就退出循环
		if len(pkFiledName)>0 {break}
	}
	return pkFiledName
}



//将miniui得到的数据进行格式化，主要是对时间进行统一，更新到单一条目
 func MiniUIDataUpdate(obj interface{}, SingleItem map[string]interface{},m orm.Params) {
	StructType := reflect.TypeOf(obj).Elem() //通过反射获取type定义

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
			}else {

				m[f.Name] = SingleItem[f.Name]
			}

		}
	}

}
