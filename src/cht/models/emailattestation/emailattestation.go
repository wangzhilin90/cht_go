package emailattestation

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/go-gomail/gomail"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CheckEmailUseRequestStruct struct {
	Email                string
	UserID               int32
	ChengHuiTongTraceLog string
}

type UserAttestationSaveStruct struct {
	UserID               int32
	EmailStatus          int32
	ChengHuiTongTraceLog string
}

type UserEmailSaveRequestStruct struct {
	Email                string
	UserID               int32
	ChengHuiTongTraceLog string
}

type SendEmailRequestStruct struct {
	UserID               int32
	SendTo               string
	Subject              string
	Content              string
	IP                   string
	Addtime              int32
	ChengHuiTongTraceLog string
}

/**
 * [CheckEmailUse 查询jl_user表,如果查询到用户说明邮箱已被使用返回1，否则返回0]
 * @DateTime 2017-09-18T15:10:09+0800
 */
func CheckEmailUse(ceurs *CheckEmailUseRequestStruct) int32 {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("CheckEmailUse input param:", ceurs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id").
		From("jl_user").
		Where(fmt.Sprintf("id != %d", ceurs.UserID)).
		And(fmt.Sprintf("email='%v'", ceurs.Email))

	sql := qb.String()
	Logger.Debug("CheckEmailUse sql:", sql)
	var id int32
	err := o.Raw(sql).QueryRow(&id)
	if err != nil {
		Logger.Debugf("CheckEmailUse query failed %v", err)
		return 0
	}
	Logger.Debugf("CheckEmailUse used")
	return 1
}

func InsertUserAttestation(uesrs *UserAttestationSaveStruct) error {
	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_user_attestation (user_id,email_status,email_passtime) values(?,?,?)")
	sql := buf.String()
	Logger.Debugf("InsertUserAttestation sql %v", sql)

	var email_passtime int64
	if uesrs.EmailStatus == 2 {
		email_passtime = time.Now().Unix()
	}

	_, err := o.Raw(sql, uesrs.UserID, uesrs.EmailStatus, email_passtime).Exec()
	if err != nil {
		Logger.Errorf("InsertUserAttestation insert failed", err)
		return err
	}
	return nil
}

func UpdateUserAttestation(uesrs *UserAttestationSaveStruct) error {
	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("update jl_user_attestation set email_status=? ,email_passtime=? where user_id=?")
	sql := buf.String()
	Logger.Debugf("UpdateUserAttestation sql: %v", sql)

	var email_passtime int64
	if uesrs.EmailStatus == 2 {
		email_passtime = time.Now().Unix()
	}

	res, _ := o.Raw(sql, uesrs.EmailStatus, email_passtime, uesrs.UserID).Exec()
	num, _ := res.RowsAffected()
	if num == 0 {
		err := fmt.Errorf("UpdateUserAttestation update failed")
		Logger.Errorf("UpdateUserAttestation update failed")
		return err
	}
	return nil
}

/**
 * [UserAttestationSave 根据user_id查询认证表看是否存在该用户的认证数据，
 * 如果存在就更新数据，否则插入数据,更新或插入成功返回1，否则返回0]
 * @DateTime 2017-09-18T15:21:54+0800
 */
func UserAttestationSave(uesrs *UserAttestationSaveStruct) int32 {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("UserAttestationSave input param:", uesrs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("user_id").
		From("jl_user_attestation").
		Where(fmt.Sprintf("user_id=%d", uesrs.UserID)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("UserAttestationSave sql:", sql)
	var user_id int32
	err := o.Raw(sql).QueryRow(&user_id)
	if err != nil {
		Logger.Debugf("UserAttestationSave query failed %v", err)
		err = InsertUserAttestation(uesrs)
		if err != nil {
			return 0
		}
		return 1
	}
	Logger.Debugf("UserAttestationSave used")
	err = UpdateUserAttestation(uesrs)
	if err != nil {
		return 0
	}
	return 1
}

//根据user_id修改用户的email，成功返回1，失败0
func UserEmailSave(uesrs *UserEmailSaveRequestStruct) int32 {
	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("update jl_user set email=? where id=?")
	sql := buf.String()
	Logger.Debugf("UserEmailSave sql: %v", sql)

	res, _ := o.Raw(sql, uesrs.Email, uesrs.UserID).Exec()
	num, _ := res.RowsAffected()
	if num == 0 {
		Logger.Errorf("UserEmailSave update failed")
		return 0
	}
	return 1
}

/**
 * [InsertEmailLog  插入jl_sendmsg表，插入数据成功得到插入ID]
 * @param    sers *SendEmailRequestStruct 请求入参
 * @return   int32	成功后返回最新一条插入的ID
 * @DateTime 2017-09-18T17:24:37+0800
 */
func InsertEmailLog(sers *SendEmailRequestStruct) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("insert into jl_sendmsg (id,user_id,send_to,subject,content,addtime,ip) values(next VALUE FOR MYCATSEQ_SENDMSG,?,?,?,?,?,?)")
	sql := buf.String()
	Logger.Debugf("InsertEmailLog sql: %v", sql)

	res, err := o.Raw(sql,
		sers.UserID,
		sers.SendTo,
		sers.Subject,
		sers.Content,
		time.Now().Unix(),
		sers.IP,
	).Exec()
	if err != nil {
		Logger.Errorf("InsertEmailLog insert failed", err)
		return 0, err
	}

	num, _ := res.LastInsertId()
	Logger.Debugf("InsertEmailLog success, last insert num %d", num)
	return int32(num), nil
}

func UpdateEmailLog(lastInsertNum int32) error {
	o := orm.NewOrm()
	o.Using("default")
	buf := bytes.Buffer{}
	buf.WriteString("update jl_sendmsg set status=1,posttime=? where id=?")
	sql := buf.String()
	Logger.Debugf("UpdateEmailLog sql: %v", sql)

	res, _ := o.Raw(sql, time.Now().Unix(), lastInsertNum).Exec()
	num, _ := res.RowsAffected()
	if num == 0 {
		err := fmt.Errorf("UpdateEmailLog update failed")
		Logger.Errorf("UpdateEmailLog update failed")
		return err
	}
	Logger.Debugf("UpdateEmailLog success")
	return nil
}

func SendSmtpMail(sers *SendEmailRequestStruct) error {
	Logger.Debug("SendSmtpMail input param:", sers)
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "service1@chenghuitong.net", "chenghuitong") // 发件人
	m.SetHeader("To", m.FormatAddress(sers.SendTo, "cht"))                  // 收件人
	m.SetHeader("Subject", sers.Subject)                                    // 主题
	m.SetBody("text/html", sers.Content)                                    // 正文

	//发送邮件服务器、端口、发件人账号、发件人密码
	d := gomail.NewPlainDialer("smtp.263.net", 25, "service2@chenghuitong.net", "ChengHuiTong@2013-")
	if err := d.DialAndSend(m); err != nil {
		Logger.Errorf("SendSmtpMail send failed:%v", err)
		return err
	}
	return nil
}

func SendEmail(sers *SendEmailRequestStruct) int32 {
	num, _ := InsertEmailLog(sers)
	err := SendSmtpMail(sers)
	if err != nil {
		return 0
	}
	UpdateEmailLog(num)
	return 1
}
