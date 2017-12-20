package usertimes

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

type UserTimesDetailsRequest struct {
	Username             string
	Isadmin              int32
	Type                 int32
	ChengHuiTongTraceLog string
}

type UserTimesDetails struct {
	Username  string `orm:"column(username)"`
	IP        string `orm:"column(ip)"`
	Logintime int32  `orm:"column(logintime)"`
	Times     int32  `orm:"column(times)"`
	Isadmin   int32  `orm:"column(isadmin)"`
	Type      int32  `orm:"column(type)"`
}

type UserTimesUpdateRequest struct {
	Username             string
	IP                   string
	Isadmin              int32
	Type                 int32
	ChengHuiTongTraceLog string
}

type UserTimesInsertRequest struct {
	Username             string
	IP                   string
	Isadmin              int32
	Type                 int32
	ChengHuiTongTraceLog string
}

type UserTimesDeleteRequest struct {
	Username             string
	Type                 int32
	ChengHuiTongTraceLog string
}

func GetUserTimesDetails(utdr *UserTimesDetailsRequest) (*UserTimesDetails, error) {
	Logger.Debugf("GetUserTimesDetails input param:%v", utdr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user_times").
		Where("1=1").
		And(fmt.Sprintf("type=%d", utdr.Type))

	if utdr.Username != "" {
		qb.And(fmt.Sprintf("username=\"%s\"", utdr.Username))
	}

	if utdr.Isadmin != 0 {
		qb.And(fmt.Sprintf("isadmin=%d", utdr.Isadmin))
	}

	qb.Limit(1)
	sql := qb.String()
	Logger.Debugf("GetUserTimesDetails sql:%v", sql)
	var utd UserTimesDetails
	err := o.Raw(sql).QueryRow(&utd)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetUserTimesDetails query failed:%v", err)
		return nil, err
	}

	Logger.Debugf("GetUserTimesDetails return value:%v", utd)
	return &utd, nil
}

func UpdateUserTimes(udur *UserTimesUpdateRequest) bool {
	Logger.Debugf("UpdateUserTimes input param:%v", udur)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Update("jl_user_times")

	var str string
	if udur.IP != "" {
		str += fmt.Sprintf("ip=\"%s\",", udur.IP)
	}

	str += fmt.Sprintf("logintime=%d,", time.Now().Unix())
	str += fmt.Sprintf("times=times+1,")

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("username=\"%s\"", udur.Username))
	qb.And(fmt.Sprintf("type=%d", udur.Type))
	if udur.Isadmin != 0 {
		qb.And(fmt.Sprintf("isadmin=%d", udur.Isadmin))
	}

	sql := qb.String()
	Logger.Debugf("UpdateUserTimes sql:%v", sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		Logger.Errorf("UpdateUserTimes update failed:%v", err)
		return false
	}

	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}

func InsertUserTimes(utir *UserTimesInsertRequest) bool {
	Logger.Debugf("InsertUserTimes input param:%v", utir)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_user_times (username,ip,logintime,times,isadmin,type) ")
	buf.WriteString("values (?,?,?,?,?,?)")

	sql := buf.String()
	last_id, err := o.Raw(sql,
		utir.Username,
		utir.IP,
		time.Now().Unix(),
		1,
		utir.Isadmin,
		utir.Type).Exec()
	if err != nil {
		Logger.Errorf("InsertUserTimes failed:%v", err)
		return false
	}
	Logger.Debugf("InsertUserTimes insert last id:%v", last_id)
	return true
}

func DeleteUserTimes(utdr *UserTimesDeleteRequest) bool {
	Logger.Debugf("DeleteUserTimes input param:%v", utdr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("DELETE FROM jl_user_times WHERE username=? AND type=?")
	sql := buf.String()

	Logger.Debugf("DeleteUserTimes sql:%v", sql)
	res, err := o.Raw(sql, utdr.Username, utdr.Type).Exec()
	if err != nil {
		Logger.Errorf("DeleteUserTimes query failed:%v", err)
		return false
	}

	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	Logger.Debugf("DeleteUserTimes affect num:%v", num)
	return true
}
