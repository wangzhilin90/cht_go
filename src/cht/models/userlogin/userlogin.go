package userlogin

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

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

type UserlLoginRequest struct {
	Username             string
	Password             string
	IP                   string
	Isadmin              int32
	ChengHuiTongTraceLog string
}

type UserInfoResult struct {
	ID       int32  `orm:"column(id)"`
	UserName string `orm:"column(username)"`
	Password string `orm:"column(password)"`
	Email    string `orm:"column(email)"`
	Islock   bool   `orm:"column(islock)"`
}

/**
 * [GetLoginFailedTimes 获取用户登录已失败次数]
 * @param    ulr *UserlLoginRequest 请求入参
 * @return	 int32 错误次数 error函数是否错误
 * @DateTime 2017-08-28T14:39:48+0800
 */
func GetLoginFailedTimes(ulr *UserlLoginRequest) (int32, error) {
	if ulr == nil || ulr.Username == "" {
		err := fmt.Errorf("input param nil")
		Logger.Errorf("input param nil")
		return 0, err
	}

	o := orm.NewOrm()
	o.Using("default")

	var temp int32
	o.Raw("SELECT times FROM jl_user_times WHERE username=?", ulr.Username).QueryRow(&temp)
	Logger.Debugf("GetLoginFailedTimes res", temp)
	return temp, nil
}

/**
 * [CheckLoginUserExists 检查错误次数表用户是否存在]
 * @param    username string 用户名
 * @DateTime 2017-08-28T14:35:03+0800
 */
func CheckLoginUserExists(ulr *UserlLoginRequest) (*UserInfoResult, bool, error) {
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return nil, false, err
	}

	o := orm.NewOrm()
	o.Using("default")

	var uir UserInfoResult
	err := o.Raw("SELECT id,username,password,email,islock FROM  jl_user where username=? or email=? or phone=?", ulr.Username, ulr.Username, ulr.Username).QueryRow(&uir)
	if err == orm.ErrNoRows {
		return nil, false, nil
	} else if err != nil {
		Logger.Error("CheckLoginUserExists query failed", err.Error())
		return nil, false, err
	}
	Logger.Debugf("GetLoginFailedTimes res", uir)
	return &uir, true, nil
}

/**
 * [Checkpassword 检查密码和数据库是否一致]
 * @param    ulr *UserlLoginRequest  用户请求入参
 * @return 	 bool true为一致，false为不一致
 * @DateTime 2017-08-28T14:42:38+0800
 */
func Checkpassword(ulr *UserlLoginRequest) bool {
	if ulr == nil || ulr.Username == "" {
		return false
	}

	res, b, err := CheckLoginUserExists(ulr)
	if err != nil || b == false || res.Password != ulr.Password {
		return false
	}
	return true
}

/**
 * [CheckUserTimesTbExist 检查错误表中是否存在]
 * @param    ulr *UserlLoginRequest 用户请求入参
 * @DateTime 2017-08-28T18:52:52+0800
 */
func CheckUserTimesTbExist(ulr *UserlLoginRequest) (bool, error) {
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return false, err
	}

	o := orm.NewOrm()
	o.Using("default")

	var userNm string
	err := o.Raw("SELECT username  FROM  jl_user_times where username=?", ulr.Username).QueryRow(&userNm)
	if err != nil {
		Logger.Debug("CheckUserTimesTbExist query failed")
		return false, nil
	}
	Logger.Debugf("GetLoginFailedTimes res", userNm)
	return true, nil
}

func InsertUserTimesTb(ulr *UserlLoginRequest) (bool, error) {
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return false, err
	}

	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("insert into jl_user_times (username,ip,logintime,times,isadmin) values(?,?,?,?,?)", ulr.Username, ulr.IP, time.Now().Unix(), 1, ulr.Isadmin).Exec()
	if err != nil {
		err := fmt.Errorf("InsertUserTimesTb failed", err)
		Logger.Error("insert mysql failed")
		return false, err
	}
	return true, nil
}

func UpdateUserTimesTb(ulr *UserlLoginRequest) (bool, error) {
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return false, err
	}

	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("update jl_user_times set ip=?,logintime=?,times=times+1", ulr.IP, time.Now().Unix()).Exec()
	if err != nil {
		err := fmt.Errorf("UpdateUserTimesTb failed", err)
		Logger.Error("update mysql failed")
		return false, err
	}
	return true, nil
}

func DeleteUserTimesTb(ulr *UserlLoginRequest) (bool, error) {
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return false, err
	}

	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Raw("delete from jl_user_times where username=?", ulr.Username).Exec()
	if err != nil {
		err := fmt.Errorf("DeleteUserTimesTb failed", err)
		Logger.Error("delete mysql failed")
		return false, err
	}
	return true, nil
}
