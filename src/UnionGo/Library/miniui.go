package Library

import (
//	"reflect"
//	"github.com/astaxie/beego/orm"
//	"encoding/json"
//	"fmt"
//	"time"
//	"strings"
)

type MiniuiGrid struct {
	Total int64 `json:"total"`
	Data  interface{} `json:"data"`
}

//构造新的struct,接收json 绑定
//	[{"_state":"modified","Guid":22,"Msgexpdate":"2014-08-15T15:56:18"},{"_state":"modified","Guid":23,"Msgexpdate":"2014-08-15T15:56:10"}]

type DataList struct {
	List [] map[string]interface {}
}
