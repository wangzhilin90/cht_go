package articlecate

import (
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ArticleCateListRequest struct {
	ID                   int32
	Name                 string
	Keywords             string
	Description          string
	Pid                  int32
	Status               int32
	ChengHuiTongTraceLog string
}

type ArticleCateDetails struct {
	ID          int32  `orm:"column(id)"`
	Name        string `orm:"column(name)"`
	Keywords    string `orm:"column(keywords)"`
	Description string `orm:"column(description)"`
	Pid         int32  `orm:"column(pid)"`
	Status      int32  `orm:"column(status)"`
	ImgURL      string `orm:"column(img_url)"`
	Sort        int32  `orm:"column(sort)"`
	Addtime     int32  `orm:"column(addtime)"`
}

func GetArticleCateList(aclr *ArticleCateListRequest) ([]ArticleCateDetails, error) {
	Logger.Debugf("GetArticleCateList input param:%v", aclr)
	o := orm.NewOrm()
	o.Using("default")
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("jl_article_cate").
		Where("1=1")

	if aclr.ID != 0 {
		qb.And(fmt.Sprintf("id=%d", aclr.ID))
	}

	if aclr.Name != "" {
		qb.And(fmt.Sprintf("name like \"%%%s%%\"", aclr.Name))
	}

	if aclr.Keywords != "" {
		qb.And(fmt.Sprintf("keywords like \"%%%s%%\"", aclr.Keywords))
	}

	if aclr.Description != "" {
		qb.And(fmt.Sprintf("description like \"%%%s%%\"", aclr.Description))
	}

	if aclr.Pid != -1 {
		qb.And(fmt.Sprintf("pid=%d", aclr.Pid))
	}

	if aclr.Status != -1 {
		qb.And(fmt.Sprintf("status=%d", aclr.Status))
	}

	sql := qb.String()
	Logger.Debugf("GetArticleCateList sql:%v", sql)
	var acd []ArticleCateDetails
	_, err := o.Raw(sql).QueryRows(&acd)
	if err != nil {
		Logger.Errorf("GetArticleCateList query failed:%v", err)
		return nil, err
	}

	Logger.Debugf("GetArticleCateList return value:%v", acd)
	return acd, nil
}
