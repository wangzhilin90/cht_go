package roledetails

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type RoleDetailsRequestStruct struct {
	RoleID               int32
	ChengHuiTongTraceLog string
}

type RoleDetailsStruct struct {
	ID          int32  `orm:column(id)`
	RoleName    string `orm:column(role_name)`
	Remark      string `orm:column(remark)`
	PowerConfig string `orm:column(power_config)`
	CreateTime  int32  `orm:column(create_time)`
}

/**
 * [GetRoleDetails 角色详情]
 * @param    rdrs *RoleDetailsRequestStruct 请求入参
 * @return   RoleDetailsStruct 角色详情结果
 * @DateTime 2017-10-10T15:25:24+0800
 */
func GetRoleDetails(rdrs *RoleDetailsRequestStruct) (*RoleDetailsStruct, error) {
	Logger.Debugf("GetRoleDetails input param: %v", rdrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT * FROM jl_sys_role WHERE id = ? LIMIT 1")
	sql := buf.String()

	Logger.Debugf("GetRoleDetails sql: %v", sql)

	var rds RoleDetailsStruct
	err := o.Raw(sql, rdrs.RoleID).QueryRow(&rds)
	if err != nil {
		Logger.Debugf("GetRoleDetails query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetRoleDetails return value %v", rds)
	return &rds, nil
}
