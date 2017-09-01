package loguserlogin

import (
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

type LogUserLoginResponse struct {
	UserID int32
}

type UserBorrowInfo struct {
	UserID     int32  `orm:"column(user_id)"`
	Id         int32  `orm:"column(id)"`
	AccountAct string `orm:"column(account_act)"`
	Addtime    int32  `orm:"column(addtime)"`
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
	if err != nil {
		Logger.Debugf("err", err)
		return nil, err
	}
	Logger.Debugf("GetBorrowInfo res", ubi)
	return &ubi, nil
}

func UpdateLogUserlLogin(lulr *LogUserlLoginRequest) (bool, error) {
	if lulr == nil {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return false, err
	}
	Logger.Debugf("UpdateLogUserlLogin input param", lulr)

	res, err := GetBorrowInfo(lulr)
	if err != nil {
		Logger.Error("GetBorrowInfo failed", err)
		return false, err
	}

	o := orm.NewOrm()
	o.Using("default")
	_, err = o.Raw("insert into jl_user_login_log (id,user_id,login_time,login_style,login_ip,tender_money,tender_time) values(next VALUE FOR MYCATSEQ_USER_LOGIN_LOG,?,?,?,?,?,?)", lulr.UserID, time.Now().Unix(), lulr.LoginStyle, lulr.LoginIP, res.AccountAct, res.Addtime).Exec()
	if err != nil {
		Logger.Error("insert mysql failed")
		return false, err
	}
	return true, nil
}
