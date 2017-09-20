package juanzeng

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type RequestStruct struct {
	UserID               int32
	Content              string
	ChengHuiTongTraceLog string
}

type MesslistResultStruct struct {
	Username  string `orm:"column(username)"`
	Avatar    string `orm:"column(avatar)"`
	Addtime   int32  `orm:"column(addtime)"`
	Content   string `orm:"column(content)"`
	Reply     string `orm:"column(reply)"`
	UpContent string `orm:"column(up_content)"`
	UpReply   string `orm:"column(up_reply)"`
}

type FundlistResultStruct struct {
	Type     int32  `orm:"column(type)"`
	Addtime  int32  `orm:"column(addtime)"`
	Username string `orm:"column(username)"`
	Money    string `orm:"column(money)"`
}

type NumlistResultStruct struct {
	Num   string `orm:"column(num)"`
	Money string `orm:"column(money)"`
}

type JuanzengResponseStruct struct {
	Messlist []*MesslistResultStruct
	Fundlist []*FundlistResultStruct
	Numlist  *NumlistResultStruct
	Tzr      string
}

func GetMesslistResult(rs *RequestStruct) ([]MesslistResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetMesslistResult input param:", rs)

	buf := bytes.Buffer{}
	buf.WriteString("SELECT u.username,u.avatar,p.addtime,p.content,p.reply,p.up_content,p.up_reply ")
	buf.WriteString("FROM  jl_juanmess p LEFT JOIN jl_user u ON u.id=p.user_id ")
	buf.WriteString("ORDER BY p.id desc ")
	sql := buf.String()
	Logger.Debugf("GetMesslistResult sql %v", sql)

	var mrs []MesslistResultStruct
	_, err := o.Raw(sql).QueryRows(&mrs)
	if err != nil {
		Logger.Errorf("GetMesslistResult query failed :%v", err)
		return nil, err
	}

	Logger.Debugf("GetMesslistResult res :%v", mrs)
	return mrs, nil
}

func GetFundlistResult(rs *RequestStruct) ([]FundlistResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetFundlistResult input param:", rs)

	buf := bytes.Buffer{}
	buf.WriteString("SELECT p.type,u.username,p.money,p.addtime ")
	buf.WriteString("FROM jl_point_juanzeng p LEFT JOIN jl_user u ON u.id=p.user_id ")
	buf.WriteString("ORDER BY p.id DESC LIMIT 100 ")
	sql := buf.String()
	Logger.Debugf("GetFundlistResult sql %v", sql)

	var frs []FundlistResultStruct
	_, err := o.Raw(sql).QueryRows(&frs)
	if err != nil {
		Logger.Errorf("GetFundlistResult query failed :%v", err)
		return nil, err
	}

	Logger.Debugf("GetFundlistResult res :%v", frs)
	return frs, nil
}

func GetNumlistResult(rs *RequestStruct) (*NumlistResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetNumlistResult input param:", rs)

	buf := bytes.Buffer{}
	buf.WriteString("select num,money from jl_juannum limit 1")
	sql := buf.String()
	Logger.Debugf("GetNumlistResult sql %v", sql)

	var nlrs NumlistResultStruct
	err := o.Raw(sql).QueryRow(&nlrs)
	if err != nil {
		Logger.Debugf("GetNumlistResult query failed :%v", err)
		return nil, nil
	}

	Logger.Debugf("GetNumlistResult res :%v", nlrs)
	return &nlrs, nil
}

func GetTotalJuanNum(rs *RequestStruct) (string, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetTotalJuanNum input param:", rs)

	buf := bytes.Buffer{}
	buf.WriteString("SELECT SUM(money) FROM jl_point_juanzeng")
	sql := buf.String()
	Logger.Debugf("GetTotalJuanNum sql %v", sql)

	var totalJuanNum string
	o.Raw(sql).QueryRow(&totalJuanNum)
	Logger.Debugf("GetTotalJuanNum res :%v", totalJuanNum)
	if totalJuanNum == "" {
		totalJuanNum = "0"
	}
	return totalJuanNum, nil
}

// func GetInfo(rs *RequestStruct) JuanzengResponseStruct {

// }

/**
 * [AddMess 添加留言]
 * @param    rs *RequestStruct 请求入参
 * @return   int32 返回最新插入的一条留言ID
 * @DateTime 2017-09-20T15:43:27+0800
 */
func AddMess(rs *RequestStruct) (int32, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetTotalJuanNum input param:", rs)

	buf := bytes.Buffer{}
	buf.WriteString("INSERT INTO jl_juanmess (id,content,user_id,addtime) VALUES (next VALUE FOR MYCATSEQ_JUANMESS,?,?,?)")
	sql := buf.String()
	Logger.Debugf("AddMess sql: %v", sql)

	res, err := o.Raw(sql,
		rs.Content,
		rs.UserID,
		time.Now().Unix(),
	).Exec()
	if err != nil {
		Logger.Errorf("AddMess insert failed :%v", err)
		return 0, err
	}
	lastInsertNum, _ := res.LastInsertId()
	Logger.Debugf("AddMess res :%v", lastInsertNum)
	return int32(lastInsertNum), nil
}
