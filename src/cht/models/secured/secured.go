package secured

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type SecuredRequestStruct struct {
	ChengHuiTongTraceLog string
}

type SecuredInfoStruct struct {
	Secured string `orm:"column(secured)"`
}

/**
 * [GetSecuredList 得到担保人信息]
 * @param    srs *SecuredRequestStruct 请求入参
 * @return   *SecuredInfoStruct 返回担保人列表
 * @DateTime 2017-09-13T11:29:32+0800
 */
func GetSecuredList(srs *SecuredRequestStruct) ([]SecuredInfoStruct, error) {
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT `secured` FROM jl_borrow WHERE  secured <> '' GROUP BY secured")
	sql := buf.String()

	Logger.Debugf("GetSecuredList sql %v", sql)
	var sis []SecuredInfoStruct
	_, err := o.Raw(sql).QueryRows(&sis)
	if err != nil {
		Logger.Debugf("GetSecuredList query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetSecuredList res %v", sis)
	return sis, nil
}
