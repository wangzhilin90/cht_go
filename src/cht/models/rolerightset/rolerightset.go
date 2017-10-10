package rolerightset

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RoleRightSetRequestStruct struct {
	RoleID               int32
	PowerConfig          string
	ChengHuiTongTraceLog string
}

/**
 * [SetRoleRight 角色权限编辑]
 * @param    RoleRightSetRequestStruct 请求入参
 * @return   bool true:权限修改成功  false:权限修改失败
 * @DateTime 2017-10-10T16:15:44+0800
 */
func SetRoleRight(rrsrs *RoleRightSetRequestStruct) bool {
	Logger.Debugf("SetRoleRight input param: %v", rrsrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("UPDATE jl_sys_role SET power_config = ? WHERE id = ?")
	sql := buf.String()

	Logger.Debugf("SetRoleRight sql: %v", sql)

	res, err := o.Raw(sql, rrsrs.PowerConfig, rrsrs.RoleID).Exec()
	if err != nil {
		Logger.Errorf("SetRoleRight update failed %v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("SetRoleRight change num  %v", num)
	return true
}
