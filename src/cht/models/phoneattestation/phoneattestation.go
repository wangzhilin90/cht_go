package phoneattestation

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CheckPhoneUseRequest struct {
	Phone                string
	ChengHuiTongTraceLog string
}

type GetUserIdByhsidRequest struct {
	Hsid                 string //徽商ID
	ChengHuiTongTraceLog string
}

type UpdatePhoneRequest struct {
	Phone                string
	UserID               int32
	ChengHuiTongTraceLog string
}

/**
 * [CheckPhoneUse 根据手机号查询jl_user表，查找手机号是否使用过]
 * @param    cpur *CheckPhoneUseRequest 请求入参
 * @return   true:使用过 fasle:没有使用过
 * @DateTime 2017-09-14T10:40:39+0800
 */
func CheckPhoneUse(cpur *CheckPhoneUseRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	var num int32
	o.Raw("select count(1) from jl_user where phone=?", cpur.Phone).QueryRow(&num)
	Logger.Debugf("CheckPhoneUse  query num %d", num)
	if num == 0 {
		return false
	}
	return true
}

/**
 * [GetUserIdByhsid 根据hsid查询jl_user表获取该用户的ID，返回ID]
 * @param    gibr *GetUserIdByhsidRequest 请求入参
 * @return   int32 返回user_id
 * @DateTime 2017-09-14T10:35:56+0800
 */
func GetUserIdByhsid(gibr *GetUserIdByhsidRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")

	var user_id int32
	err := o.Raw("select id from jl_user where hsid=?", gibr.Hsid).QueryRow(&user_id)
	if err != nil {
		Logger.Errorf("GetUserIdByhsid query failed", err)
		return 0, err
	}
	Logger.Debugf("GetUserIdByhsid res", user_id)

	return user_id, nil
}

/**
 * [UpdatePhone 更新jl_user表修改该用户的手机号]
 * @param    upr *UpdatePhoneRequest 请求入参
 * @return   string 1000:更新手机号成功 1001:查找原手机号失败 1002：新手机号与原手机号一致 1003:插入日志失败 1004:更新手机号失败
 * @DateTime 2017-09-14T10:34:50+0800
 */
func UpdatePhone(upr *UpdatePhoneRequest) string {
	o := orm.NewOrm()
	o.Using("default")

	o.Begin()
	buf := bytes.Buffer{}
	buf.WriteString("select phone from jl_user where id=?")
	sql := buf.String()
	Logger.Debugf("UpdatePhone query oldphone sql:%v", sql)

	var oldPhone string
	err := o.Raw(sql, upr.UserID).QueryRow(&oldPhone)
	if err != nil {
		o.Rollback()
		Logger.Debug("UpdatePhone query oldphone failed")
		return "1001"
	}

	if oldPhone == upr.Phone {
		o.Rollback()
		Logger.Debug("UpdatePhone newphone same with oldphone")
		return "1002"
	}

	buf = bytes.Buffer{}
	buf.WriteString("insert into jl_phone_log (id,user_id,addtime,old_phone,new_phone) values(next VALUE FOR MYCATSEQ_PHONE_LOG,?,?,?,?)")
	sql = buf.String()
	Logger.Debugf("UpdatePhone insert sql :%v", sql)
	_, err = o.Raw(sql, upr.UserID, time.Now().Unix(), oldPhone, upr.Phone).Exec()
	if err != nil {
		o.Rollback()
		Logger.Debug("UpdatePhone insert sql failed")
		return "1003"
	}

	buf = bytes.Buffer{}
	buf.WriteString("update jl_user set phone=? where id=?")
	sql = buf.String()
	Logger.Debugf("UpdatePhone update sql :%v", sql)
	res, _ := o.Raw(sql, upr.Phone, upr.UserID).Exec()
	num, _ := res.RowsAffected()
	if num == 0 {
		o.Rollback()
		Logger.Debug("UpdatePhone update phone failed")
		return "1004"
	}
	o.Commit()
	Logger.Debugf("UpdatePhone update success ")
	return "1000"
}
