package advertmanage

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AdvertManageRequest struct {
	Type                 int32
	Limit                int32
	ChengHuiTongTraceLog string
}

type AdvertManageStruct struct {
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

func GetAdvertManage(amr *AdvertManageRequest) ([]AdvertManageStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAdvertManage input param:", amr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id,type,img,adverturl,title,addtime,adduser,status,fid,starttime,endtime").
		From("jl_advert_manage").
		Where(fmt.Sprintf("type=%d", amr.Type)).
		And(fmt.Sprintf("starttime<=%d", time.Now().Unix())).
		And(fmt.Sprintf("%d<=endtime", time.Now().Unix())).
		OrderBy("addtime").Desc()

	if amr.Limit != 0 {
		qb.Limit(int(amr.Limit))
	}
	sql := qb.String()
	Logger.Debug("GetAdvertManage sql:", sql)
	var ams []AdvertManageStruct
	_, err := o.Raw(sql).QueryRows(&ams)
	if err != nil {
		Logger.Errorf("GetAdvertManage query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetAdvertManage res :%v", ams)
	return ams, nil
}
