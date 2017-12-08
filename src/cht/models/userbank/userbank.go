package userbank

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type UserBankDetailsRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

type UserBankDetailsStruct struct {
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

type UserBankUpdateRequest struct {
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

type UserBankInsertRequest struct {
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

type UserBankCountRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

func GetUserBankDetails(ubdr *UserBankDetailsRequest) (*UserBankDetailsStruct, error) {
	Logger.Info("GetUserBankDetails input param:", ubdr)
	o := orm.NewOrm()
	o.Using("default")

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user_bank").
		Where(fmt.Sprintf("user_id=%d", ubdr.UserID)).
		Limit(1)

	sql := qb.String()
	Logger.Info("GetUserBankDetails sql:", sql)
	var ubds UserBankDetailsStruct
	err := o.Raw(sql).QueryRow(&ubds)
	if err != nil {
		Logger.Errorf("GetUserBankDetails query failed:%v", err)
		return nil, err
	}
	Logger.Info("GetUserBankDetails return value:", ubds)
	return &ubds, nil
}

func UpdateUserBank(ubur *UserBankUpdateRequest) bool {
	Logger.Info("UpdateUserBank input param:", ubur)
	o := orm.NewOrm()
	o.Using("default")

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user_bank")

	var str string
	if ubur.Name != "" {
		str += fmt.Sprintf("name=\"%s\",", ubur.Name)
	}

	if ubur.Account != "" {
		str += fmt.Sprintf("account=\"%s\",", ubur.Account)
	}

	if ubur.Branch != "" {
		str += fmt.Sprintf("branch=\"%s\",", ubur.Branch)
	}

	if ubur.Addip != "" {
		str += fmt.Sprintf("addip=\"%s\",", ubur.Addip)
	}

	if ubur.ID != 0 {
		str += fmt.Sprintf("id=%d,", ubur.ID)
	}

	if ubur.Bank != 0 {
		str += fmt.Sprintf("bank=%d,", ubur.Bank)
	}

	if ubur.Province != 0 {
		str += fmt.Sprintf("province=%d,", ubur.Province)
	}

	if ubur.City != 0 {
		str += fmt.Sprintf("city=%d,", ubur.City)
	}

	if ubur.Area != 0 {
		str += fmt.Sprintf("area=%d,", ubur.Area)
	}

	if ubur.Addtime != 0 {
		str += fmt.Sprintf("addtime=%d,", ubur.Addtime)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("user_id=%d", ubur.UserID))
	sql := qb.String()
	Logger.Info("UpdateUserBank sql:", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Errorf("UpdateUserBank update failed:%v", err)
		return false
	}
	affectNum, _ := res.RowsAffected()
	if affectNum == 0 {
		return false
	}
	return true
}

func InsertUserBank(ubir *UserBankInsertRequest) bool {
	Logger.Info("InsertUserBank input param:", ubir)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_user_bank ")
	buf.WriteString("(id,user_id,name,account,bank,branch,province,city,area,addtime,addip) ")
	buf.WriteString("values(next VALUE FOR MYCATSEQ_USER_BANK,?,?,?,?,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("InsertUserBank sql:%v", sql)

	res, err := o.Raw(sql,
		ubir.UserID,
		ubir.Name,
		ubir.Account,
		ubir.Bank,
		ubir.Branch,
		ubir.Province,
		ubir.City,
		ubir.Area,
		time.Now().Unix(),
		ubir.Addip,
	).Exec()
	if err != nil {
		Logger.Errorf("InsertUserBank insert failed:%v", err)
		return false
	}
	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("InsertUserBank rows effect num:%v", num)
	return true
}

func GetUserBankNum(ubcr *UserBankCountRequest) (int32, error) {
	Logger.Info("GetUserBankNum input param:", ubcr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT COUNT(1) AS total_num FROM jl_user_bank UB LEFT JOIN jl_user U ON UB.user_id=U.id WHERE U.id = ? LIMIT 1 ")
	sql := buf.String()
	Logger.Debugf("GetUserBankNum sql:%v", sql)

	var num int32
	err := o.Raw(sql, ubcr.UserID).QueryRow(&num)
	if err != nil {
		Logger.Errorf("GetUserBankNum query num failed:%v", err)
		return 0, err
	}

	return num, err
}
