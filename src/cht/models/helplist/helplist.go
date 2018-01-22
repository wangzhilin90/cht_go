package helplist

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type HelpListRequest struct {
	Status               int32
	Cateid               int32
	LimitOffset          int32
	LimitNum             int32
	ChengHuiTongTraceLog string
}

type HelpListResultStruct struct {
	ID      int32  `orm:"column(id)"`
	Title   string `orm:"column(title)"`
	Content string `orm:"column(content)"`
}

/*获取帮助中心文章列表*/
func GetHelpList(hr *HelpListRequest) ([]HelpListResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetHelpList input param:", hr)

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("A.id,A.title,A.content FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id").
		Where("1=1").
		And(fmt.Sprintf("A.status=%d", hr.Status)).
		And(fmt.Sprintf("A.cateid=%d", hr.Cateid))

	qb.OrderBy("A.id ASC")
	if hr.LimitNum != 0 {
		qb.Limit(int(hr.LimitNum))
	}

	if hr.LimitNum != 0 && hr.LimitOffset != 0 {
		qb.Offset(int(hr.LimitOffset))
	}
	sql := qb.String()
	Logger.Debugf("GetHelpList sql %v", sql)

	var hlrs []HelpListResultStruct
	_, err := o.Raw(sql).QueryRows(&hlrs)
	if err != nil {
		Logger.Errorf("GetHelpList query failed :%v", err)
		return nil, err
	}

	Logger.Debugf("GetHelpList res :%v", hlrs)
	return hlrs, nil
}
