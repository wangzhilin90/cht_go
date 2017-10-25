package advertupdate

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type AdvertUpdateRequest struct {
	ID                   int32
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
 * [UpdateAdvert 广告图片管理---修改广告图片]
 * @param    aup *AdvertUpdateRequest 请求入参
 * bool [description]
 * @DateTime 2017-10-25T15:25:48+0800
 */
func UpdateAdvert(aup *AdvertUpdateRequest) bool {
	Logger.Debugf("UpdateAdvert input param:%v", aup)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_advert_manage")

	var str string
	if aup.Img != "" {
		str += fmt.Sprintf("img=\"%s\",", aup.Img)
	}

	if aup.Adverturl != "" {
		str += fmt.Sprintf("adverturl=\"%s\",", aup.Adverturl)
	}

	if aup.Title != "" {
		str += fmt.Sprintf("title=\"%s\",", aup.Title)
	}

	if aup.Type != 0 {
		str += fmt.Sprintf("type=%d,", aup.Type)
	}

	if aup.Adduser != 0 {
		str += fmt.Sprintf("adduser=%d,", aup.Adduser)
	}

	if aup.Fid != 0 {
		str += fmt.Sprintf("fid=%d,", aup.Fid)
	}

	if aup.Starttime != 0 {
		str += fmt.Sprintf("starttime=%d,", aup.Starttime)
	}

	if aup.Endtime != 0 {
		str += fmt.Sprintf("endtime=%d,", aup.Endtime)
	}

	//每次修改需要更新系统当前时间
	str += fmt.Sprintf("addtime=%d,", time.Now().Unix())

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("id=%d", aup.ID))
	sql := qb.String()

	Logger.Debug("UpdateAdvert sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Debugf("UpdateAdvert update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("UpdateAdvert change num :%v", num)
	if num == 0 {
		return false
	}
	return true
}
