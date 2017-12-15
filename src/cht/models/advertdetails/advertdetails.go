package advertdetails

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AdvertDetailsRequest struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

type AdvertDetails struct {
	ID        int32  `orm:"column(id)"`
	Type      int32  `orm:"column(type)"`
	Img       string `orm:"column(img)"`
	Adverturl string `orm:"column(adverturl)"`
	Title     string `orm:"column(title)"`
	Addtime   int32  `orm:"column(addtime)"`
	Adduser   int32  `orm:"column(adduser)"`
	Status    int32  `orm:"column(status)"`
	Fid       int32  `orm:"column(fid)"`
	Starttime int32  `orm:"column(starttime)"`
	Endtime   int32  `orm:"column(endtime)"`
}

func GetAdvertDetails(adr *AdvertDetailsRequest) (*AdvertDetails, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAdvertDetails input param:", adr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_advert_manage").
		Where(fmt.Sprintf("id=%d", adr.ID)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetAdvertDetails sql:", sql)

	var ad AdvertDetails
	err := o.Raw(sql).QueryRow(&ad)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetAdvertDetails query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetAdvertDetails res %v", ad)
	return &ad, nil
}
