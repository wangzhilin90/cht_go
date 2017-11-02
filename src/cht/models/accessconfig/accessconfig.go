package accessconfig

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AccessConfigRequest struct {
	Source               string
	ChengHuiTongTraceLog string
}

type AccessConfigStruct struct {
	ID      int32  `orm:"column(id)"`
	Name    string `orm:"column(name)"`
	Source  string `orm:"column(source)"`
	Addtime int32  `orm:"column(addtime)"`
}

/**
 * [GetAccessConfig 获取推广名称记录表]
 * @param    acr *AccessConfigRequest 请求入参
 * @return   *AccessConfigStruct 推广名列表详情
 * @DateTime 2017-10-30T16:28:26+0800
 */
func GetAccessConfig(acr *AccessConfigRequest) (*AccessConfigStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAccessConfig input param:", acr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_access_config").
		Where(fmt.Sprintf("source=\"%s\"", acr.Source)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetAccessConfig sql:", sql)

	var acs AccessConfigStruct
	err := o.Raw(sql).QueryRow(&acs)
	if err != nil {
		Logger.Debugf("GetAccessConfig query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetAccessConfig res %v", acs)
	return &acs, nil
}
