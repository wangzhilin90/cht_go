package kefudutydelete

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type KefuDutyDeleteRequest struct {
	Idstr                string
	ChengHuiTongTraceLog string
}

/**
 * [DeleteKefuDuty description]
 * @param    kddr *KefuDutyDeleteRequest 请求入参
 * @return   bool true:删除成功 false:删除失败
 * @DateTime 2017-10-27T17:04:39+0800
 */
func DeleteKefuDuty(kddr *KefuDutyDeleteRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("DeleteKefuDuty input param:", kddr)

	users := strings.Split(kddr.Idstr, ",")
	var str string
	for range users {
		str += fmt.Sprintf(",%v", "?")
	}
	str = strings.TrimPrefix(str, ",")
	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM  jl_customer_plan WHERE id  IN (")
	buf.WriteString(str)
	buf.WriteString(")")
	sql := buf.String()
	Logger.Debugf("DeleteKefuDuty sql:%v", sql)

	res, err := o.Raw(sql, users).Exec()
	if err != nil {
		Logger.Errorf("DeleteKefuDuty delete failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("DeleteKefuDuty rows effect num:%v", num)
	return true
}
