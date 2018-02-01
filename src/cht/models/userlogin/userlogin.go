package userlogin

import (
	"bytes"
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
	Type                 int32
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
	Logger.Debugf("GetLoginFailedTimes input param:%v", ulr)
	if ulr == nil || ulr.Username == "" {
		err := fmt.Errorf("input param nil")
		Logger.Errorf("input param nil")
		return 0, err
	}

	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("SELECT times FROM jl_user_times WHERE username=?")
	sql := buf.String()
	Logger.Debugf("GetLoginFailedTimes sql:%v", sql)

	var temp int32
	err := o.Raw(sql, ulr.Username).QueryRow(&temp)
	if err == orm.ErrNoRows {
		return 0, nil
	} else if err != nil {
		Logger.Errorf("GetLoginFailedTimes query failed:%v", err)
		return 0, err
	}
	Logger.Debugf("GetLoginFailedTimes res", temp)
	return temp, nil
}

/**
 * [CheckLoginUserExists 检查错误次数表用户是否存在]
 * @param    username string 用户名
 * @DateTime 2017-08-28T14:35:03+0800
 */
func CheckLoginUserExists(ulr *UserlLoginRequest) (*UserInfoResult, error) {
	Logger.Debugf("CheckLoginUserExists input param:%v", ulr)
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param nil")
		err := fmt.Errorf("input param nil")
		return nil, err
	}

	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT id,username,password,email,islock FROM  jl_user where username=? or email=? or phone=?")
	sql := buf.String()
	Logger.Debugf("CheckLoginUserExists sql:%v", sql)

	var uir UserInfoResult
	err := o.Raw(sql, ulr.Username, ulr.Username, ulr.Username).QueryRow(&uir)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Error("CheckLoginUserExists query failed", err)
		return nil, err
	}
	Logger.Debugf("CheckLoginUserExists res", uir)
	return &uir, nil
}

/**
 * [Checkpassword 检查密码和数据库是否一致]
 * @param    ulr *UserlLoginRequest  用户请求入参
 * @return 	 bool true为一致，false为不一致
 * @DateTime 2017-08-28T14:42:38+0800
 */
func Checkpassword(ulr *UserlLoginRequest) bool {
	Logger.Debugf("Checkpassword input param:%v", ulr)
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("Checkpassword input not valid")
		return false
	}

	res, err := CheckLoginUserExists(ulr)
	if err != nil || res == nil || res.Password != ulr.Password {
		return false
	}
	return true
}

/**
 * [CheckUserTimesTbExist 检查错误表中是否存在]
 * @param    ulr *UserlLoginRequest 用户请求入参
 * @DateTime 2017-08-28T18:52:52+0800
 */
func CheckUserTimesTbExist(ulr *UserlLoginRequest) bool {
	Logger.Debugf("CheckUserTimesTbExist input param:%v", ulr)
	if ulr == nil || ulr.Username == "" {
		Logger.Errorf("input param not valid")
		return false
	}

	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("SELECT username  FROM  jl_user_times where username=?")
	sql := buf.String()
	Logger.Debugf("CheckUserTimesTbExist sql:%v", sql)

	var userNm string
	err := o.Raw(sql, ulr.Username).QueryRow(&userNm)
	if err == orm.ErrNoRows {
		Logger.Debugf("CheckUserTimesTbExist query nil")
		return false
	} else if err != nil {
		Logger.Errorf("CheckUserTimesTbExist query failed:%v", err)
		return false
	}
	Logger.Debugf("CheckUserTimesTbExist res", userNm)
	return true
}

func InsertUserTimesTb(ulr *UserlLoginRequest) (bool, error) {
	Logger.Debugf("InsertUserTimesTb input param:%v", ulr)
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
	Logger.Debugf("UpdateUserTimesTb input param:%v", ulr)
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
	Logger.Debugf("DeleteUserTimesTb input param:%v", ulr)
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

func InsertLoginLog(ulr *UserlLoginRequest) bool {
	Logger.Debugf("InsertLoginLog input param:%v", ulr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_login_log (id,account,addtime,type,ip,window,status) values ")
	buf.WriteString("(next VALUE FOR MYCATSEQ_LOGIN_LOG,?,?,?,?,?,?) ")
	sql := buf.String()

	Logger.Debugf("InsertLoginLog sql:%v", sql)
	res, err := o.Raw(sql,
		ulr.Username,
		time.Now().Unix(),
		ulr.Type,
		ulr.IP,
		1, 0).Exec()

	if err != nil {
		Logger.Errorf("InsertLoginLog failed:%v", err)
		return false
	}

	num, _ := res.RowsAffected()
	if num == 0 {
		return false
	}
	return true
}
