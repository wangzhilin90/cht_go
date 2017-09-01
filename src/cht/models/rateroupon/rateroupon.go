package rateroupon

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type CouponTable struct {
	ID           int32  `orm:"column(id)"`
	UserID       int32  `orm:"column(user_id)"`
	Addtime      int32  `orm:"column(addtime)"`
	StartTime    int32  `orm:"column(start_time)"`
	EndTime      int32  `orm:"column(end_time)"`
	UseTime      int32  `orm:"column(use_time)"`
	Status       int32  `orm:"column(status)"`
	TenderID     int32  `orm:"column(tender_id)"`
	Apr          string `orm:"column(apr)"`
	AppAdd       string `orm:"column(app_add)"`
	MinTender    string `orm:"column(min_tender)"`
	MaxTender    string `orm:"column(max_tender)"`
	TimeLimit    string `orm:"column(time_limit)"`
	BorrowType   string `orm:"column(borrow_type)"`
	Name         string `orm:"column(name)"`
	Remark       string `orm:"column(remark)"`
	ActivityName string `orm:"column(activity_name)"`
}

type CouponRequest struct {
	UserID               int32
	Status               int32
	Limit                int32
	ChengHuiTongTraceLog string
	OrderBy              string
}

// func init() {
// 	orm.Debug = true
// 	orm.RegisterModel(new(CouponTable))
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
// 	user := "cht"
// 	passwd := "cht123456"
// 	host := "192.168.10.2"
// 	port := "3306"
// 	dbname := "chtlocal"
// 	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, dbname))
// }

func GetRateRoupon(crs *CouponRequest) (cps []CouponTable, err error) {
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_apr_coupon").
		Where(fmt.Sprintf("user_id=%d", crs.UserID))

	if crs.Status != 0 {
		qb.And(fmt.Sprintf("status=%d", crs.Status))
	}
	if crs.OrderBy != "" {
		qb.OrderBy(crs.OrderBy)
	}
	if crs.Limit != 0 {
		qb.Limit(int(crs.Limit))
	}

	sql := qb.String()
	Logger.Debugf("GetRateRoupon sql", sql)
	var cs []CouponTable
	o.Raw(sql).QueryRows(&cs)
	Logger.Debugf("GetRateRoupon res", cs)
	return cs, nil
}
