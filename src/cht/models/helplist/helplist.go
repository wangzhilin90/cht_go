package helplist

import (
	"bytes"
	. "cht/common/logger"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type HelpListRequest struct {
	Status               int32
	Cateid               int32
	ChengHuiTongTraceLog string
}

type HelpListResultStruct struct {
	Title   string `orm:column(title)`
	Content string `orm:column(content)`
}

/*获取帮助中心文章列表*/
func GetHelpList(hr *HelpListRequest) ([]HelpListResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetHelpList input param:", hr)

	buf := bytes.Buffer{}
	buf.WriteString("SELECT A.title,A.content FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id ")
	buf.WriteString("WHERE A.status=? AND A.cateid=? ORDER BY A.id ASC ")
	sql := buf.String()
	Logger.Debugf("GetHelpList sql %v", sql)

	var hlrs []HelpListResultStruct
	_, err := o.Raw(sql, hr.Status, hr.Cateid).QueryRows(&hlrs)
	if err != nil {
		Logger.Errorf("GetHelpList query failed :%v", err)
		return nil, err
	}

	Logger.Debugf("GetHelpList res :%v", hlrs)
	return hlrs, nil
}
