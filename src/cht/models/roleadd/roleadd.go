package roleadd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RoleAddRequestStruct struct {
	RoleName             string
	Remark               string
	ChengHuiTongTraceLog string
}

/**
 * [AddRole 添加角色]
 * @param    rars *RoleAddRequestStruct 请求入参
 * @return   bool  true:表示添加成功 ，false:表示添加失败
 * @DateTime 2017-10-10T14:51:24+0800
 */
func AddRole(rars *RoleAddRequestStruct) bool {
	Logger.Debugf("AddRole input param: %v", rars)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("INSERT INTO jl_sys_role (id,role_name,remark,create_time,power_config) VALUES (next VALUE FOR MYCATSEQ_SYS_ROLE,?,?,?,?)")
	sql := buf.String()

	Logger.Debugf("AddRole sql: %v", sql)

	res, err := o.Raw(sql, rars.RoleName, rars.Remark, time.Now().Unix(), "0").Exec()
	if err != nil {
		Logger.Errorf("AddRole insert failed %v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("AddRole change num  %v", num)
	return true
}
