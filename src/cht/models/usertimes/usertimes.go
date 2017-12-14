package usertimes

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type UserTimesDetailsRequest struct {
	Username             string
	Isadmin              int32
	ChengHuiTongTraceLog string
}

type UserTimesDetails struct {
	Username  string `orm:"column(username)"`
	IP        string `orm:"column(ip)"`
	Logintime int32  `orm:"column(logintime)"`
	Times     int32  `orm:"column(times)"`
	Isadmin   int32  `orm:"column(isadmin)"`
}

type UserTimesUpdateRequest struct {
	Username             string `thrift:"username,1" db:"username" json:"username"`
	IP                   string `thrift:"ip,2" db:"ip" json:"ip"`
	Logintime            int32  `thrift:"logintime,3" db:"logintime" json:"logintime"`
	Times                int32  `thrift:"times,4" db:"times" json:"times"`
	Isadmin              int32  `thrift:"isadmin,5" db:"isadmin" json:"isadmin"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,6" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

type UserTimesInsertRequest struct {
	Username             string `thrift:"username,1" db:"username" json:"username"`
	IP                   string `thrift:"ip,2" db:"ip" json:"ip"`
	Logintime            int32  `thrift:"logintime,3" db:"logintime" json:"logintime"`
	Times                int32  `thrift:"times,4" db:"times" json:"times"`
	Isadmin              int32  `thrift:"isadmin,5" db:"isadmin" json:"isadmin"`
	ChengHuiTongTraceLog string `thrift:"chengHuiTongTraceLog,6" db:"chengHuiTongTraceLog" json:"chengHuiTongTraceLog"`
}

func GetUserTimesDetails(utdr *UserTimesDetailsRequest) (*UserTimesDetails, error) {
	Logger.Debugf("GetUserTimesDetails input param:%v", utdr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user_times").
		Where("1=1")

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

	if udur.Logintime != 0 {
		str += fmt.Sprintf("logintime=%d,", udur.Logintime)
	}

	if udur.Times != 0 {
		str += fmt.Sprintf("times=%d,", udur.Times)
	}

	if udur.Isadmin != 0 {
		str += fmt.Sprintf("isadmin=%d,", udur.Isadmin)
	}

	str = strings.TrimSuffix(str, ",")
	qb.Set(str)
	qb.Where(fmt.Sprintf("username=\"%s\"", udur.Username))
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
	buf.WriteString("insert into jl_user_times (username,ip,logintime,times,isadmin) ")
	buf.WriteString("values (?,?,?,?,?)")

	sql := buf.String()
	last_id, err := o.Raw(sql,
		utir.Username,
		utir.IP,
		utir.Logintime,
		utir.Times,
		utir.Isadmin).Exec()
	if err != nil {
		Logger.Errorf("InsertUserTimes failed:%v", err)
		return false
	}
	Logger.Debugf("InsertUserTimes insert last id:%v", last_id)
	return true
}
