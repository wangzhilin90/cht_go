package advertdel

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type AdvertDelRequest struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

/**
 * [GetAdvertFid 获取广告图片地址]
 * @param    adr *AdvertDelRequest 请求入参
 * @return   int32 图片地址
 * @DateTime 2017-10-25T11:56:50+0800
 */
func GetAdvertFid(adr *AdvertDelRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetAdvertFid input param:", adr)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("fid").
		From("jl_advert_manage").
		Where(fmt.Sprintf("id=%d", adr.ID)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetAdvertFid sql:", sql)

	var fid int32
	err := o.Raw(sql).QueryRow(&fid)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetAdvertFid query failed %v", err)
		return 0, err
	}
	Logger.Debugf("GetAdvertFid res %v", fid)
	return fid, nil
}

/**
 * [DelAdvert 广告图片管理---删除广告图片]
 * @param    adr *AdvertDelRequest 请求入参
 * @return   bool true:删除成功  false:删除失败
 * @DateTime 2017-10-25T11:43:42+0800
 */
func DelAdvert(adr *AdvertDelRequest) bool {
	Logger.Debug("DelAdvert input param:", adr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM  jl_advert_manage WHERE id = ?")
	sql := buf.String()
	Logger.Debugf("DelAdvert sql:%v", sql)

	res, err := o.Raw(sql, adr.ID).Exec()
	if err != nil {
		Logger.Errorf("DelAdvert delete failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("DelAdvert rows effect num:%v", num)
	return true
}
