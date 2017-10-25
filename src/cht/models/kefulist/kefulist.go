package kefulist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type KeFuListRequestStruct struct {
	RoleID               int32
	Status               int32
	CustomerType         string
	ChengHuiTongTraceLog string
}

type KeFuDetailsStruct struct {
	ID            int32  `orm:"column(id)"`
	RoleID        int32  `orm:"column(role_id)"`
	Account       string `orm:"column(account)"`
	Realname      string `orm:"column(realname)"`
	Password      string `orm:"column(password)"`
	Mobile        string `orm:"column(mobile)"`
	Qq            string `orm:"column(qq)"`
	Lastloginip   string `orm:"column(lastloginip)"`
	Lastlogintime int32  `orm:"column(lastlogintime)"`
	CreateTime    int32  `orm:"column(create_time)"`
	Status        int32  `orm:"column(status)"`
	Views         int32  `orm:"column(views)"`
	CustomerType  int32  `orm:"column(customer_type)"`
}

func GetKeFuList(kfrs *KeFuListRequestStruct) ([]KeFuDetailsStruct, error) {
	Logger.Debugf("GetKeFuList input param:%v", *kfrs)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_sys_user").
		Where("1=1")

	if kfrs.Status != 0 {
		qb.And(fmt.Sprintf("status=%d", kfrs.Status))
	}

	if kfrs.RoleID != 0 {
		qb.And(fmt.Sprintf("role_id=%d", kfrs.RoleID))
	}

	if kfrs.CustomerType != "" {
		qb.And(fmt.Sprintf("customer_type IN (%v)", kfrs.CustomerType))
	}

	sql := qb.String()
	Logger.Debugf("GetKeFuList sql:%v", sql)
	var kfds []KeFuDetailsStruct
	_, err := o.Raw(sql).QueryRows(&kfds)
	if err != nil {
		Logger.Errorf("GetKeFuList query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetKeFuList res:%v", kfds)
	return kfds, nil
}
