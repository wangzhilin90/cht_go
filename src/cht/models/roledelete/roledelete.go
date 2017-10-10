package roledelete

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type RoleDeleteRequestStruct struct {
	RoleIDStr            string
	ChengHuiTongTraceLog string
}

/**
 * [DeleteRole 角色删除]
 * @param    rdrs *RoleDeleteRequestStruct 请求入参
 * @return   bool true:删除成功 fasle:删除失败
 * @DateTime 2017-10-10T15:18:30+0800
 */
func DeleteRole(rdrs *RoleDeleteRequestStruct) bool {
	Logger.Debugf("DeleteRole input param: %v", rdrs)
	o := orm.NewOrm()
	o.Using("default")

	rold_ID_arr := strings.Split(rdrs.RoleIDStr, ",")
	var str string
	for range rold_ID_arr {
		str += fmt.Sprintf(",%v", "?")
	}
	str = strings.TrimPrefix(str, ",")
	Logger.Debugf("str %v", str)

	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM jl_sys_role WHERE id !=1 AND id IN (")
	buf.WriteString(str)
	buf.WriteString(")")
	sql := buf.String()

	Logger.Debugf("DeleteRole sql: %v", sql)
	res, err := o.Raw(sql, rold_ID_arr).Exec()
	if err != nil {
		Logger.Errorf("DeleteRole delete failed %v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("DeleteRole change num  %v", num)
	return true
}
