package advertadd

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AdvertAddRequest struct {
	Type                 int32
	Img                  string
	Adverturl            string
	Title                string
	Adduser              int32
	Fid                  int32
	Starttime            int32
	Endtime              int32
	ChengHuiTongTraceLog string
}

/**
 * [AddAdvert 添加广告图片]
 * @param    aar *AdvertAddRequest 请求入参
 * @return   bool true:添加成功  false:添加失败
 * @DateTime 2017-10-25T10:26:36+0800
 */
func AddAdvert(aar *AdvertAddRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	Logger.Debug("AddAdvert input param:", aar)
	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_advert_manage ")
	buf.WriteString("(id,type,img,adverturl,title,addtime,adduser,fid,starttime,endtime) ")
	buf.WriteString("values(next VALUE FOR MYCATSEQ_ADVERT_MANAGE,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("AddAdvert sql:%v", sql)

	res, err := o.Raw(sql,
		aar.Type,
		aar.Img,
		aar.Adduser,
		aar.Title,
		time.Now().Unix(),
		aar.Adduser,
		aar.Fid,
		aar.Starttime,
		aar.Endtime,
	).Exec()
	if err != nil {
		Logger.Errorf("AddAdvert insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("AddAdvert rows effect num:%v", num)
	return true
}
