package sysconfig

import (
	// "bytes"
	. "cht/common/logger"
	"cht/utils"
	// "fmt"
	// "github.com/astaxie/beego"
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
	var scs []SysConfigStruct

	err := utils.GetCache("sys:config", &scs)
	if err != nil {
		// cache_expire, _ := beego.AppConfig.Int("cache_expire")
		// Logger.Debugf("cache_expire:%v", cache_expire)
		o := orm.NewOrm()
		o.Using("default")
		_, err = o.Raw("select  * from jl_sys_config").QueryRows(&scs)
		if err != nil {
			Logger.Errorf("GetSysConfig query failed: %v", err)
			return nil, err
		}
		utils.SetCache("sys:config", scs, 60)
	}

	Logger.Debugf("GetSysConfig res %v", scs)
	return scs, nil
}
