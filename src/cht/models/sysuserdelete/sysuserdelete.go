package sysuserdelete

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type SysUserDeleteRequest struct {
	UserIDStr            string
	ChengHuiTongTraceLog string
}

/**
 * [DeleteSysUser 删除后台管理用户]
 * @param    sudr *SysUserDeleteRequest 请求入参
 * @return   bool true:删除成功  false:删除失败
 * @DateTime 2017-10-17T17:03:19+0800
 */
func DeleteSysUser(sudr *SysUserDeleteRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("DeleteSysUser input param:", sudr)

	users := strings.Split(sudr.UserIDStr, ",")
	var str string
	for range users {
		str += fmt.Sprintf(",%v", "?")
	}
	str = strings.TrimPrefix(str, ",")
	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM  jl_sys_user WHERE id  IN (")
	buf.WriteString(str)
	buf.WriteString(")")
	sql := buf.String()
	Logger.Debugf("DeleteSysUser sql:%v", sql)

	res, err := o.Raw(sql, users).Exec()
	if err != nil {
		Logger.Errorf("DeleteSysUser delete failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("DeleteSysUser rows effect num:%v", num)
	return true
}
