package sysuseradd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SysUserAddRequest struct {
	Account              string
	Password             string
	Realname             string
	Mobile               string
	Qq                   string
	Status               int32
	RoleID               int32
	CustomerType         int32
	CreateTime           int32
	Lastlogintime        int32
	Views                int32
	Lastloginip          string
	ChengHuiTongTraceLog string
}

/**
 * [addSysUser 添加后台管理用户]
 * @param    suars *SysUserAddRequestStruct 请求入参
 * @return   bool   true:添加成功 false:添加失败
 * @DateTime 2017-10-17T09:37:42+0800
 */
func AddSysUser(suars *SysUserAddRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("AddSysUser input param:", suars)
	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_sys_user ")
	buf.WriteString("(id,account,password,realname,mobile,qq,status,role_id,customer_type,create_time,lastlogintime,views,lastloginip) ")
	buf.WriteString("values(next VALUE FOR MYCATSEQ_SYS_USER,?,?,?,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("AddSysUser sql:%v", sql)

	res, err := o.Raw(sql,
		suars.Account,
		suars.Password,
		suars.Realname,
		suars.Mobile,
		suars.Qq,
		suars.Status,
		suars.RoleID,
		suars.CustomerType,
		suars.CreateTime,
		suars.Lastlogintime,
		suars.Views,
		suars.Lastloginip,
	).Exec()
	if err != nil {
		Logger.Errorf("AddSysUser insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("rows effect num:%v", num)
	return true
}
