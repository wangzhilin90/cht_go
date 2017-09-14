package phoneattestation

import (
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
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
 * @return   更新成功返回true，否则返回false
 * @DateTime 2017-09-14T10:34:50+0800
 */
func UpdatePhone(upr *UpdatePhoneRequest) bool {
	o := orm.NewOrm()
	o.Using("default")

	res, _ := o.Raw("update jl_user set phone=? where id=?", upr.Phone, upr.UserID).Exec()
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
