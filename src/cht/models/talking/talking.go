package talking

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TalkingRequest struct {
	Cateid               int32
	Status               int32
	ChengHuiTongTraceLog string
}

type TalkListResult struct {
	ID      int32  `orm:"column(id)"`
	Title   string `orm:"column(title)"`
	ImgURL  string `orm:"column(img_url)"`
	Content string `orm:"column(content)"`
}

func GetTalkinglist(tr *TalkingRequest) ([]TalkListResult, error) {
	Logger.Debugf("GetTalkinglist input param:%v", tr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT A.id,A.title,A.img_url,A.content FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id ")
	buf.WriteString("WHERE A.cateid=? AND A.status=? AND A.title LIKE '%交流会%' ORDER BY A.id DESC LIMIT 4")
	sql := buf.String()
	Logger.Debugf("GetTalkinglist sql: %v", sql)

	var tlr []TalkListResult
	_, err := o.Raw(sql, tr.Cateid, tr.Status).QueryRows(&tlr)
	if err != nil {
		Logger.Errorf("GetTalkinglist query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetTalkinglist res %v", tlr)
	return tlr, nil
}

func GetOnelist(tr *TalkingRequest) ([]TalkListResult, error) {
	Logger.Debugf("GetOnelist input param:%v", tr)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("SELECT A.id,A.title,A.img_url,A.content FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id ")
	buf.WriteString("WHERE A.id=3760 ORDER BY A.sort,A.addtime desc")
	sql := buf.String()
	Logger.Debugf("GetOnelist sql: %v", sql)

	var tlr []TalkListResult
	_, err := o.Raw(sql).QueryRows(&tlr)
	if err != nil {
		Logger.Errorf("GetOnelist query failed %v", err)
		return nil, err
	}
	Logger.Debugf("GetOnelist res %v", tlr)
	return tlr, nil
}
