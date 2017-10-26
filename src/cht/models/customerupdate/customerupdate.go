package customerupdate

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type CustomerUpdateRequest struct {
	ID                   string `thrift:"id,1" db:"id" json:"id"`
	Islock               int32  `thrift:"islock,2" db:"islock" json:"islock"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,3" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

/**
 * [UpdateCustomer 更新用户表是否锁定]
 * @param    cur *CustomerUpdateRequest 请求入参
 * @return   bool true:更新成功 false:更新失败
 * @DateTime 2017-10-26T15:33:58+0800
 */
func UpdateCustomer(cur *CustomerUpdateRequest) bool {
	Logger.Debugf("UpdateCustomer input param:%v", cur)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user").
		Set(fmt.Sprintf("islock=?")).
		Where(fmt.Sprintf("id = ?"))

	sql := qb.String()
	Logger.Debug("UpdateCustomer sql:", sql)

	//用于一次 prepare 多次 exec，以提高批量执行的速度
	p, _ := o.Raw(sql).Prepare()
	userIdStr := strings.Split(cur.ID, ",")
	for _, userID := range userIdStr {
		_, err := p.Exec(cur.Islock, userID)
		if err != nil {
			Logger.Errorf("UpdateCustomer userID %v, udpate failed:%v", userID, err)
			return false
		}
	}
	return true
}
