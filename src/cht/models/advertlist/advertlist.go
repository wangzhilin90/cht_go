package advertlist

import (
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AdvertListRequest struct {
	ChengHuiTongTraceLog string
}

type AdvertList struct {
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

/**
 * [GetAdvertListTatalNum 广告图片管理---列表总数]
 * @param    alr *AdvertListRequest 请求入参
 * @return   int32 总数
 * @DateTime 2017-10-25T14:50:45+0800
 */
func GetAdvertListTatalNum(alr *AdvertListRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAdvertListTatalNum input param:", alr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("count(*)").
		From("jl_advert_manage")

	sql := qb.String()
	Logger.Debug("GetAdvertListTatalNum sql:", sql)

	var totalNum int32
	err := o.Raw(sql).QueryRow(&totalNum)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetAdvertListTatalNum query failed:%v", err)
		return 0, err
	}
	Logger.Debugf("GetAdvertListTatalNum res:%v", totalNum)
	return totalNum, nil
}

/**
 * [GetAdvertList 广告图片管理---列表服务]
 * @param    alr *AdvertListRequest 请求入参
 * @return   []AdvertList 广告图片列表
 * @DateTime 2017-10-25T14:49:57+0800
 */
func GetAdvertList(alr *AdvertListRequest) ([]AdvertList, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAdvertList input param:", alr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_advert_manage")

	sql := qb.String()
	Logger.Debug("GetAdvertList sql:", sql)

	var al []AdvertList
	_, err := o.Raw(sql).QueryRows(&al)
	if err != nil {
		Logger.Errorf("GetAdvertList query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetAdvertList res:%v", al)
	return al, nil
}
