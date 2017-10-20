package userattestionbaseinfosave

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type UserAttestionBaseInfoSaveRequest struct {
	UserID               int32  `orm:"column(user_id)"`
	VideoPic             string `orm:"column(video_pic)"`
	RealStatus           int32  `orm:"column(real_status)"`
	EmailStatus          int32  `orm:"column(email_status)"`
	PhoneStatus          int32  `orm:"column(phone_status)"`
	VideoStatus          int32  `orm:"column(video_status)"`
	SceneStatus          int32  `orm:"column(scene_status)"`
	RealPasstime         int32  `orm:"column(real_passtime)"`
	ChengHuiTongTraceLog string `orm:"column(chengHuiTongTraceLog)"`
}

/**
 * [SaveUserAttestionBaseInfo 保存用户认证信息]
 * @param    uabsr *UserAttestionBaseInfoSaveRequest 请求入参
 * @return   bool  true:保存成功，false:保存失败
 * @DateTime 2017-10-19T16:49:31+0800
 */
func SaveUserAttestionBaseInfo(uabsr *UserAttestionBaseInfoSaveRequest) bool {
	Logger.Debugf("SaveUserAttestionBaseInfo input param:%v", uabsr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user_attestation")

	var str string
	if strings.TrimSpace(uabsr.VideoPic) != "" {
		str += fmt.Sprintf("video_pic=\"%s\",", uabsr.VideoPic)
	}

	if uabsr.RealStatus != 0 {
		str += fmt.Sprintf("real_status=%d,", uabsr.RealStatus)
	}

	if uabsr.EmailStatus != 0 {
		str += fmt.Sprintf("email_status=%d,", uabsr.EmailStatus)
	}

	if uabsr.PhoneStatus != 0 {
		str += fmt.Sprintf("phone_status=%d,", uabsr.PhoneStatus)
	}

	if uabsr.VideoStatus != 0 {
		str += fmt.Sprintf("video_status=%d,", uabsr.VideoStatus)
	}

	if uabsr.SceneStatus != 0 {
		str += fmt.Sprintf("scene_status=%d,", uabsr.SceneStatus)
	}

	if uabsr.RealPasstime != 0 {
		str += fmt.Sprintf("real_passtime=%d,", uabsr.RealPasstime)
	}
	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("user_id=%d", uabsr.UserID))
	sql := qb.String()

	Logger.Debug("SaveUserAttestionBaseInfo sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Debugf("SaveUserAttestionBaseInfo update failed :%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	Logger.Debugf("SaveUserAttestionBaseInfo change num :%v", num)
	if num == 0 {
		return false
	}
	return true
}
