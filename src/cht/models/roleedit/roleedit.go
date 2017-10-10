package roleedit

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RoleEditRequestStruct struct {
	RoleID               int32
	RoleName             string
	Remark               string
	ChengHuiTongTraceLog string
}

/**
 * [EditRole 角色编辑]
 * @param     rers *RoleEditRequestStruct 请求入参
 * @return    bool  true:编辑成功  false:编辑失败
 * @DateTime 2017-10-10T15:55:56+0800
 */
func EditRole(rers *RoleEditRequestStruct) bool {
	Logger.Debugf("EditRole input param: %v", rers)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("UPDATE jl_sys_role SET role_name = ?,remark = ?,create_time = ?  WHERE id = ?")
	sql := buf.String()

	Logger.Debugf("EditRole sql: %v", sql)

	res, err := o.Raw(sql, rers.RoleName, rers.Remark, time.Now().Unix(), rers.RoleID).Exec()
	if err != nil {
		Logger.Errorf("EditRole update failed %v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("EditRole change num  %v", num)
	return true
}
