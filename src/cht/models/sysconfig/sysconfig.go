package sysconfig

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SysConfigRequestStruct struct {
	ChengHuiTongTraceLog string
}

type SysConfigStruct struct {
	ID    int32  `orm:"column(id)"`
	Nid   string `orm:"column(nid)"`
	Value string `orm:"column(value)"`
	Name  string `orm:"column(name)"`
}

/**
 * [GetSysConfig 系统所有配置文件]
 * @param    []SysConfigStruct 返回系统配置项
 * @DateTime 2017-09-19T14:58:14+0800
 */
func GetSysConfig(scrs *SysConfigRequestStruct) ([]SysConfigStruct, error) {
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("select  * from jl_sys_config")
	sql := buf.String()
	Logger.Debugf("GetSysConfig sql: %v", sql)

	var scs []SysConfigStruct
	_, err := o.Raw(sql).QueryRows(&scs)
	if err != nil {
		Logger.Debugf("GetSysConfig query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetSysConfig res %v", scs)
	return scs, nil
}