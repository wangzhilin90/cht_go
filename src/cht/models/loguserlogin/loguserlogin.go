package loguserlogin

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type LogUserlLoginRequest struct {
	UserID               int32  `orm:"user_id"`
	LoginIP              string `orm:"login_ip"`
	LoginStyle           int32  `orm:"login_style"`
	ChengHuiTongTraceLog string `orm:"chengHuiTongTraceLog"`
}

type UserBorrowInfo struct {
	UserID     int32  `orm:"column(user_id)"`
	Id         int32  `orm:"column(id)"`
	AccountAct string `orm:"column(account_act)"`
	Addtime    int32  `orm:"column(addtime)"`
}

type UserLoginLogDetailsRequest struct {
	UserID               int32
	ChengHuiTongTraceLog string
}

type UserLoginLogDetails struct {
	ID          int32  `orm:"column(id)"`
	UserID      int32  `orm:"column(user_id)"`
	LoginTime   int32  `orm:"column(login_time)"`
	LoginStyle  int32  `orm:"column(login_style)"`
	LoginIP     string `orm:"column(login_ip)"`
	TenderMoney string `orm:"column(tender_money)"`
	TenderTime  int32  `orm:"column(tender_time)"`
}

// func init() {
// 	orm.Debug = true
// 	orm.RegisterDriver("mysql", orm.DRMySQL)
// 	user := "cht"
// 	passwd := "cht123456"
// 	host := "192.168.10.2"
// 	port := "3306"
// 	dbname := "chtlocal"
// 	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, passwd, host, port, dbname))
// }

func GetBorrowInfo(lulr *LogUserlLoginRequest) (*UserBorrowInfo, error) {
	if lulr == nil {
		err := fmt.Errorf("input param nil")
		Logger.Errorf("input param nil")
		return nil, err
	}
	o := orm.NewOrm()
	o.Using("default")
	var ubi UserBorrowInfo
	err := o.Raw("SELECT user_id,id,account_act,addtime FROM  jl_borrow_tender where user_id=? order by id desc limit 1", lulr.UserID).QueryRow(&ubi)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetBorrowInfo query failed:%v", err)
		return nil, err
	}
	return &ubi, nil
}

func UpdateLogUserlLogin(lulr *LogUserlLoginRequest) (bool, error) {
	Logger.Debugf("UpdateLogUserlLogin input param", lulr)
	res, err := GetBorrowInfo(lulr)
	if err != nil {
		Logger.Error("GetBorrowInfo failed:", err)
		return false, err
	}
	Logger.Debug("GetBorrowInfo res:", res)
	var accountAct string
	var addTime int64
	if res != nil {
		accountAct = res.AccountAct
		addTime = int64(res.Addtime)
	} else {
		accountAct = "0.00"
		addTime = 0
	}

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_user_login_log (id,user_id,login_time,login_style,login_ip,tender_money,tender_time) values(next VALUE FOR MYCATSEQ_USER_LOGIN_LOG,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("UpdateLogUserlLogin sql:%v", sql)

	o := orm.NewOrm()
	o.Using("default")

	_, err = o.Raw(sql,
		lulr.UserID,
		time.Now().Unix(),
		lulr.LoginStyle,
		lulr.LoginIP,
		accountAct,
		addTime).Exec()
	if err != nil {
		Logger.Error("insert mysql failed", err)
		return false, err
	}
	return true, nil
}

func GetUserLoginLogDetails(ulldr *UserLoginLogDetailsRequest) (*UserLoginLogDetails, error) {
	Logger.Debugf("GetUserLoginLogDetails input param:%v", ulldr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_user_login_log").
		Where("1=1")

	if ulldr.UserID != 0 {
		qb.And(fmt.Sprintf("user_id=%d", ulldr.UserID))
	}
	qb.OrderBy("id").Desc().Limit(1)

	sql := qb.String()
	Logger.Debugf("GetUserLoginLogDetails sql:%v", sql)
	var ulld UserLoginLogDetails
	err := o.Raw(sql).QueryRow(&ulld)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetUserLoginLogDetails query failed:%v", err)
		return nil, err
	}

	Logger.Debugf("GetUserLoginLogDetails return value:%v", ulld)
	return &ulld, nil
}
