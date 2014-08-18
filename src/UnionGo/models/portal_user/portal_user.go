package portal_user

import (
	"github.com/astaxie/beego/orm"
	"time"
	. "UnionGo/Library"
)

type Portal_user struct{
	Guid                   int
	User_code              string `orm:"pk" form:"User_code"`
	User_name              string `form:"User_name"`
	Dpt_name               string
	Msgexpdate             time.Time
}

func init() {
	orm.RegisterModel(new(Portal_user))
}

func (h Portal_user) SaveList(data string) {

	SaveMiniuiData(h,data)


}


