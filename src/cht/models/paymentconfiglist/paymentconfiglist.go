package paymentconfiglist

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentConfigListRequest struct {
	ChengHuiTongTraceLog string
}

type PaymentConfigDetails struct {
	ID     int32  `orm:"column(id)"`
	Type   int32  `orm:"column(type)"`
	Nid    string `orm:"column(nid)"`
	Name   string `orm:"column(name)"`
	Logo   string `orm:"column(logo)"`
	Config string `orm:"column(config)"`
	Fee    string `orm:"column(fee)"`
	Status int32  `orm:"column(status)"`
	Remark string `orm:"column(remark)"`
	Sort   int32  `orm:"column(sort)"`
}

/**
 * [GetPaymentConfigList 第三方支付方式列表]
 * @param    pclr *PaymentConfigListRequest 请求入参
 * @return   []PaymentConfigDetails 支付列表详情
 * @DateTime 2017-10-27T14:16:31+0800
 */
func GetPaymentConfigList(pclr *PaymentConfigListRequest) ([]PaymentConfigDetails, error) {
	Logger.Debug("GetPaymentConfigList input param:", pclr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("select * from jl_payment_config order by sort")
	sql := buf.String()
	Logger.Debugf("GetPaymentConfigList sql %v", sql)

	var pcd []PaymentConfigDetails
	_, err := o.Raw(sql).QueryRows(&pcd)
	if err != nil {
		Logger.Errorf("GetPaymentConfigList query failed :%v", err)
		return nil, err
	}
	Logger.Debugf("GetPaymentConfigList return value:%v", pcd)
	return pcd, nil
}
