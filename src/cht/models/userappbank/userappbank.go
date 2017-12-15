package userappbank

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type UserAppBankDetailsRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

type UserAppBankDetailsStruct struct {
	ID       int32  `orm:"column(id)"`
	UserID   int32  `orm:"column(user_id)"`
	Name     string `orm:"column(name)"`
	Account  string `orm:"column(account)"`
	Bank     int32  `orm:"column(bank)"`
	Branch   string `orm:"column(branch)"`
	Province int32  `orm:"column(province)"`
	City     int32  `orm:"column(city)"`
	Area     int32  `orm:"column(area)"`
	Addtime  int32  `orm:"column(addtime)"`
	Addip    string `orm:"column(addip)"`
}

//  - ChengHuiTongTraceLog
type UserAppBankUpdateRequest struct {
	ID                   int32
	UserID               int32
	Name                 string
	Account              string
	Bank                 int32
	Branch               string
	Province             int32
	City                 int32
	Area                 int32
	Addtime              int32
	Addip                string
	ChengHuiTongTraceLog string
}

type UserAppBankInsertRequest struct {
	ID                   int32
	UserID               int32
	Name                 string
	Account              string
	Bank                 int32
	Branch               string
	Province             int32
	City                 int32
	Area                 int32
	Addtime              int32
	Addip                string
	ChengHuiTongTraceLog string
}

type UserAppBankDeleteRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

/**
 * [GetUserAppBankDetails 得到app会员信息详情]
 * @param    {[type]}                 uabdr *UserAppBankDetailsRequest) ([]UserAppBankDetailsStruct, error [description]
 * @DateTime 2017-12-06T13:37:20+0800
 */
func GetUserAppBankDetails(uabdr *UserAppBankDetailsRequest) (*UserAppBankDetailsStruct, error) {
	Logger.Debugf("GetUserAppBankDetails input param:", uabdr)
	o := orm.NewOrm()
	o.Using("default")

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user_appbank").
		Where(fmt.Sprintf("user_id=%d", uabdr.UserID)).
		Limit(1)

	sql := qb.String()
	Logger.Info("GetUserAppBankDetails sql:", sql)
	var uabds UserAppBankDetailsStruct
	err := o.Raw(sql).QueryRow(&uabds)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetUserAppBankDetails query failed:%v", err)
		return nil, err
	}
	Logger.Debugf("GetUserAppBankDetails return value:", uabds)
	return &uabds, nil
}

func UpdateUserAppBank(uabur *UserAppBankUpdateRequest) bool {
	Logger.Info("UpdateUserAppBank input param:", uabur)
	o := orm.NewOrm()
	o.Using("default")

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user_appbank")

	var str string
	if uabur.Name != "" {
		str += fmt.Sprintf("name=\"%s\",", uabur.Name)
	}

	if uabur.Account != "" {
		str += fmt.Sprintf("account=\"%s\",", uabur.Account)
	}

	if uabur.Branch != "" {
		str += fmt.Sprintf("branch=\"%s\",", uabur.Branch)
	}

	if uabur.Addip != "" {
		str += fmt.Sprintf("addip=\"%s\",", uabur.Addip)
	}

	if uabur.ID != 0 {
		str += fmt.Sprintf("borrow_name=%d,", uabur.ID)
	}

	if uabur.Bank != 0 {
		str += fmt.Sprintf("bank=%d,", uabur.Bank)
	}

	if uabur.Province != 0 {
		str += fmt.Sprintf("province=%d,", uabur.Province)
	}

	if uabur.City != 0 {
		str += fmt.Sprintf("city=%d,", uabur.City)
	}

	if uabur.Area != 0 {
		str += fmt.Sprintf("area=%d,", uabur.Area)
	}

	if uabur.Addtime != 0 {
		str += fmt.Sprintf("addtime=%d,", uabur.Addtime)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("user_id=%d", uabur.UserID))
	sql := qb.String()
	Logger.Info("UpdateUserAppBank sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Errorf("UpdateUserAppBank update failed:%v", err)
		return false
	}
	affectNum, _ := res.RowsAffected()
	if affectNum == 0 {
		return false
	}
	return true
}

func InsertUserAppBank(iabir *UserAppBankInsertRequest) bool {
	Logger.Info("InsertUserAppBank input param:", iabir)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_user_appbank ")
	buf.WriteString("(id,user_id,name,account,bank,branch,province,city,area,addtime,addip) ")
	buf.WriteString("values(next VALUE FOR MYCATSEQ_USER_APPBANK,?,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("InsertUserAppBank sql:%v", sql)

	res, err := o.Raw(sql,
		iabir.UserID,
		iabir.Name,
		iabir.Account,
		iabir.Bank,
		iabir.Branch,
		iabir.Province,
		iabir.City,
		iabir.Area,
		time.Now().Unix(),
		iabir.Addip,
	).Exec()
	if err != nil {
		Logger.Errorf("InsertUserAppBank insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("InsertUserAppBank rows effect num:%v", num)
	return true
}

func DeletetUserAppBank(uabdr *UserAppBankDeleteRequest) bool {
	Logger.Info("DeletetUserAppBank input param:", uabdr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM jl_user_appbank WHERE user_id = ?")
	sql := buf.String()
	Logger.Info("DelAdvert sql:", sql)

	res, err := o.Raw(sql, uabdr.UserID).Exec()
	if err != nil {
		Logger.Errorf("DeletetUserAppBank delete failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("DeletetUserAppBank rows effect num:%v", num)
	return true
}
