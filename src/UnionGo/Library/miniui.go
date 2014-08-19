package Library

import ("reflect"
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



func save2DB(vt reflect.Type,v map[string]interface {}){
	o := orm.NewOrm()

	m := make(orm.Params)
	for i := 0; i < vt.NumField(); i++ {
		f := vt.Field(i)
		//fmt.Println(f.Name, f.Type, reflect.TypeOf( v[f.Name]))
		if v[f.Name] != nil {
			if f.Type == reflect.TypeOf(time.Now()) {
				//对时间格式进行特殊的处理，进行时区转换，miniui过来的json默认为+08:00
				//处理为go转换string为时间需要的标准时间格式
				ss := fmt.Sprintf("%s", v[f.Name])
				ss = strings.Replace(ss, "T", " ", -1)
				ss = strings.Replace(ss, "+08:00", " +08:00", -1)
				t, _ := time.Parse("2006-01-02 15:04:05 -07:00 ", ss)
				m[f.Name] = t
			}else {
				m[f.Name] = v[f.Name]
			}

		}
		//fmt.Println(f.Name, v[f.Name])
	}

	num, err := o.QueryTable(vt.Name()).Filter("Guid", v["Guid"]).Update(m)

	fmt.Println("Affected Num: %d, %s", num, err)
}

func SaveMiniuiData(h interface{}, data string) {

	//整理为可识别格式
	var s Data
	json.Unmarshal([]byte(data), &s)


	//按struct 遍历得到定义，及得到的值
	//	var h Portal_user
	vt := reflect.TypeOf(h)

	for _, v := range s.List {
		if state := v["_state"]; state!=nil {
			if state.(string)=="modified"{
			}
		}
		//fmt.Println(k, v["_state"], v)
		save2DB(vt,v)

	}

}

