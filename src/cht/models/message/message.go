package message

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MessageRequest struct {
	Smsid                int32  //短信id
	Phone                string //手机号
	Addtime              string //添加时间
	Type                 int32  //类型
	ChengHuiTongTraceLog string
}

type MessageInfoStruct struct {
	ID         int32  `orm:"column(id)"`
	Type       int32  `orm:column(type)`
	UserID     int32  `orm:column(user_id)`
	SendTo     string `orm:column(send_to)`
	Subject    string `orm:column(subject)`
	Content    string `orm:column(content)`
	Attachment string `orm:column(attachment)`
	Addtime    int32  `orm:column(addtime)`
	IP         string `orm:column(ip)`
	Posttime   int32  `orm:column(posttime)`
	Status     int32  `orm:column(status)`
}

/**
 * [GetMessageInfo 获取短信详情]
 * @param    mr *MessageRequest 请求入参
 * @return   MessageInfoStruct  返回短信详情
 * @DateTime 2017-09-11T17:09:41+0800
 */
func GetMessageInfo(mr *MessageRequest) (*MessageInfoStruct, error) {
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("select  * from jl_sendmsg ")
	buf.WriteString("where id =?  AND send_to=? limit 1")
	sql := buf.String()

	Logger.Debugf("GetMessageInfo sql %v", sql)
	var mis MessageInfoStruct
	err := o.Raw(sql, mr.Smsid, mr.Phone).QueryRow(&mis)
	if err != nil {
		Logger.Debugf("GetMessageInfo query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetMessageInfo res %v", mis)
	return &mis, nil
}

/**
 * [GetMessageCount 获取短信记录数]
 * @param    mr *MessageRequest 请求入参
 * @return   int32	返回短信记录数
 * @DateTime 2017-09-11T17:09:21+0800
 */
func GetMessageCount(mr *MessageRequest) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("select  COUNT(1)  from  jl_sendmsg SM ")
	buf.WriteString("LEFT JOIN jl_user U ON SM.user_id=U.id ")
	buf.WriteString("where SM.addtime=? and SM.send_to=? and type=1")
	sql := buf.String()

	Logger.Debugf("GetMessageCount sql %v", sql)

	var MessageNum int32
	err := o.Raw(sql, mr.Addtime, mr.Phone).QueryRow(&MessageNum)
	if err != nil {
		Logger.Error("GetMessageCount  query failed:", err)
		return 0, err
	}

	Logger.Debugf("GetMessageCount res %v", MessageNum)
	return MessageNum, nil
}
