package articledetails

import (
	"bytes"
	. "cht/common/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type ArticleDetailsRequestStruct struct {
	ID                   int32
	ChengHuiTongTraceLog string
}

type ArticleDetailsResultStruct struct {
	ID          int32  `orm:"column(id)"`
	Cateid      int32  `orm:"column(cateid)"`
	Title       string `orm:"column(title)"`
	Content     string `orm:"column(content)"`
	Keywords    string `orm:"column(keywords)"`
	Description string `orm:"column(description)"`
	ImgURL      string `orm:"column(img_url)"`
	Sort        string `orm:"column(sort)"`
	Status      int32  `orm:"column(status)"`
	Addtime     int32  `orm:"column(addtime)"`
	BannerURL   string `orm:"column(banner_url)"`
	Isbanner    int32  `orm:"column(isbanner)"`
	Type        int32  `orm:"column(type)"`
	Name        string `orm:"column(name)"`
}

type NextRequestStruct struct {
	ID                   int32
	Cateid               int32
	Type                 int32
	Addtime              int32
	ChengHuiTongTraceLog string
}

/*获取指定文章内容详情*/
func GetArticleDetails(adrs *ArticleDetailsRequestStruct) (*ArticleDetailsResultStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetArticleDetails input param:", adrs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("A.*,AC.name").
		From("jl_article A").
		LeftJoin("jl_article_cate AC").
		On("A.cateid=AC.id").
		Where(fmt.Sprintf("A.id=%d", adrs.ID)).
		Limit(1)

	sql := qb.String()
	Logger.Debug("GetArticleDetails sql:", sql)
	var adrst ArticleDetailsResultStruct
	err := o.Raw(sql).QueryRow(&adrst)
	if err != nil {
		Logger.Debugf("GetArticleDetails query failed %v", err)
		return nil, nil
	}
	Logger.Debugf("GetArticleDetails res %v", adrst)
	return &adrst, nil
}

/**
 * [UpdateReadNum 更新阅读量]
 * @param    adrs *ArticleDetailsRequestStruct 请求入参
 * @return   int32 返回受影响的行数
 * @DateTime 2017-10-12T16:41:56+0800
 */
func UpdateReadNum(adrs *ArticleDetailsRequestStruct) (int32, error) {
	Logger.Debugf("UpdateReadNum input param:%v", adrs)
	o := orm.NewOrm()
	o.Using("default")

	buf := bytes.Buffer{}
	buf.WriteString("UPDATE `jl_article` SET `read_num` = `read_num`+'1' ")
	buf.WriteString("WHERE id= ? ")
	sql := buf.String()
	Logger.Debugf("UpdateReadNum sql: %v", sql)

	res, err := o.Raw(sql, adrs.ID).Exec()
	if err != nil {
		Logger.Errorf("UpdateReadNum query failed %v", err)
		return 0, err
	}
	effectNum, _ := res.RowsAffected()
	Logger.Debugf("UpdateReadNum effect num: %v", effectNum)
	return int32(effectNum), nil
}

/*上一篇文章*/
func GetPrevArticle(nrs *NextRequestStruct) (*ArticleDetailsResultStruct, error) {
	Logger.Debugf("PrevArticle input param:%v", nrs)
	o := orm.NewOrm()
	o.Using("default")

	if nrs.Cateid == 10 {
		//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.addtime<{addtime} ORDER BY A.addtime desc
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("A.*,AC.name").
			From("jl_article A").
			LeftJoin("jl_article_cate AC").
			On("A.cateid=AC.id").
			Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
			And("A.status=1").
			And(fmt.Sprintf("A.addtime<%d", nrs.Addtime)).
			OrderBy("A.addtime").
			Desc().
			Limit(1)

		sql := qb.String()
		Logger.Debugf("PrevArticle sql: %v", sql)
		var adrs ArticleDetailsResultStruct
		err := o.Raw(sql).QueryRow(&adrs)
		if err != nil {
			Logger.Debugf("PrevArticle query failed %v", err)
			return nil, nil
		}
		Logger.Debugf("PrevArticle return value: %v", adrs)
		return &adrs, nil
	} else {
		//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.id<3195 AND type=1 ORDER BY A.id desc
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("A.*,AC.name").
			From("jl_article A").
			LeftJoin("jl_article_cate AC").
			On("A.cateid=AC.id").
			Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
			And(fmt.Sprintf("A.status=1")).
			And(fmt.Sprintf("A.id<%d", nrs.ID)).
			And(fmt.Sprintf("type=%d", nrs.Type)).
			OrderBy("A.id").Desc().
			Limit(1)

		sql := qb.String()
		Logger.Debugf("PrevArticle sql: %v", sql)
		var adrs ArticleDetailsResultStruct
		err := o.Raw(sql).QueryRow(&adrs)
		if err != nil {
			Logger.Debugf("PrevArticle query failed %v", err)
			return nil, nil
		}
		Logger.Debugf("PrevArticle return value: %v", adrs)
		return &adrs, nil
	}
}

/*下一篇文章*/
func GetNextArticle(nrs *NextRequestStruct) (*ArticleDetailsResultStruct, error) {
	Logger.Debugf("NextArticle input param:%v", nrs)
	o := orm.NewOrm()
	o.Using("default")

	if nrs.Cateid == 10 {
		//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.addtime>{addtime}
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("A.*,AC.name").
			From("jl_article A").
			LeftJoin("jl_article_cate AC").
			On("A.cateid=AC.id").
			Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
			And("A.status=1").
			And(fmt.Sprintf("A.addtime>%d", nrs.Addtime)).
			Limit(1)
		sql := qb.String()
		Logger.Debugf("NextArticle sql: %v", sql)
		var adrs ArticleDetailsResultStruct
		err := o.Raw(sql).QueryRow(&adrs)
		if err != nil {
			Logger.Debugf("NextArticle query failed %v", err)
			return nil, nil
		}
		Logger.Debugf("NextArticle return value: %v", adrs)
		return &adrs, nil
	} else {
		//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.id>3195 AND type=1
		qb, _ := orm.NewQueryBuilder("mysql")
		qb.Select("A.*,AC.name").
			From("jl_article A").
			LeftJoin("jl_article_cate AC").
			On("A.cateid=AC.id").
			Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
			And(fmt.Sprintf("A.status=1")).
			And(fmt.Sprintf("A.id>%d", nrs.ID)).
			And(fmt.Sprintf("type=%d", nrs.Type)).
			Limit(1)
		sql := qb.String()
		Logger.Debugf("NextArticle sql: %v", sql)
		var adrs ArticleDetailsResultStruct
		err := o.Raw(sql).QueryRow(&adrs)
		if err != nil {
			Logger.Debugf("NextArticle query failed %v", err)
			return nil, nil
		}
		Logger.Debugf("NextArticle return value: %v", adrs)
		return &adrs, nil
	}
}
