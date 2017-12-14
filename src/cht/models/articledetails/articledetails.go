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
	Status               int32 //默认不传为-1，不用拼接；为0时也需要拼接jl_article表
	ChengHuiTongTraceLog string
}

type ArticleDetailsStruct struct {
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
	Sort                 int32
	Prefix               string
	IsApp                int32 //请求上、下页接口是先判断此参数,如果is_app=1则为app端请求，请按下述APP文章上、下页逻辑判断
	ChengHuiTongTraceLog string
}

/*获取指定文章内容详情*/
func GetArticleDetails(adrs *ArticleDetailsRequestStruct) (*ArticleDetailsStruct, error) {
	o := orm.NewOrm()
	o.Using("default")
	Logger.Debug("GetArticleDetails input param:", adrs)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("A.*,AC.name").
		From("jl_article A").
		LeftJoin("jl_article_cate AC").
		On("A.cateid=AC.id").
		Where(fmt.Sprintf("A.id=%d", adrs.ID))

	if adrs.Status != -1 {
		qb.And(fmt.Sprintf("A.status=%d", adrs.Status))
	}

	qb.Limit(1)
	sql := qb.String()
	Logger.Debug("GetArticleDetails sql:", sql)
	var adrst ArticleDetailsStruct
	err := o.Raw(sql).QueryRow(&adrst)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("GetArticleDetails query failed %v", err)
		return nil, err
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
func GetPrevArticle(nrs *NextRequestStruct) (*ArticleDetailsStruct, error) {
	Logger.Debugf("PrevArticle input param:%v", nrs)
	o := orm.NewOrm()
	o.Using("default")
	var qb, _ = orm.NewQueryBuilder("mysql")
	//pc端逻辑
	if nrs.IsApp == 0 {
		if nrs.Cateid == 10 {
			//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.addtime<{addtime} ORDER BY A.addtime desc
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
		} else {
			//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.id<3195 AND type=1 ORDER BY A.id desc
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
		}
	} else if nrs.IsApp == 1 { //app端逻辑
		switch nrs.Cateid {
		case 4:
			//$prevWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND addtime>{$addtime}) OR sort<{$sort})";
			//$prevOrder = 'sort DESC, addtime ASC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("((A.sort=%d AND A.addtime > %d) OR A.sort < %d)", nrs.Sort, nrs.Addtime, nrs.Sort)).
				OrderBy("sort DESC, addtime ASC").
				Limit(1)
		case 5:
			//$prevWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND id>{$id}) OR sort<{$sort})";
			//$prevOrder = 'sort DESC, id ASC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("((A.sort=%d AND A.id > %d) OR A.sort < %d)", nrs.Sort, nrs.ID, nrs.Sort)).
				OrderBy("sort DESC, id ASC").
				Limit(1)
		case 8:
			//$prevWhere = "status=1 AND cateid={$cateid} AND title LIKE '{$prefix}%' AND ((sort={$sort} AND id>{$id}) OR sort<{$sort})";
			//$prevOrder = 'sort DESC, id ASC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("A.title LIKE \"%s%%\"", nrs.Prefix)).
				And(fmt.Sprintf("((A.sort=%d AND A.id > %d) OR A.sort < %d)", nrs.Sort, nrs.ID, nrs.Sort)).
				OrderBy("sort DESC, id ASC").
				Limit(1)
		}
	}
	sql := qb.String()
	Logger.Debugf("PrevArticle sql: %v", sql)
	var adrs ArticleDetailsStruct
	err := o.Raw(sql).QueryRow(&adrs)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("PrevArticle query failed %v", err)
		return nil, err
	}
	Logger.Debugf("PrevArticle return value: %v", adrs)
	return &adrs, nil
}

/*下一篇文章*/
func GetNextArticle(nrs *NextRequestStruct) (*ArticleDetailsStruct, error) {
	Logger.Debugf("NextArticle input param:%v", nrs)
	o := orm.NewOrm()
	o.Using("default")
	var qb, _ = orm.NewQueryBuilder("mysql")

	if nrs.IsApp == 0 {
		if nrs.Cateid == 10 {
			//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.addtime>{addtime}
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And("A.status=1").
				And(fmt.Sprintf("A.addtime>%d", nrs.Addtime)).
				Limit(1)
		} else {
			//SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.id>3195 AND type=1
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.id>%d", nrs.ID)).
				And(fmt.Sprintf("type=%d", nrs.Type)).
				Limit(1)
		}
	} else if nrs.IsApp == 1 {
		switch nrs.Cateid {
		case 4:
			//$nextWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND addtime<{$addtime}) OR sort>{$sort})";
			//$nextOrder = 'sort ASC, addtime DESC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("((A.sort=%d AND A.addtime < %d) OR A.sort > %d)", nrs.Sort, nrs.Addtime, nrs.Sort)).
				OrderBy("sort ASC, addtime DESC").
				Limit(1)
		case 5:
			//$nextWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND id<{$id}) OR sort>{$sort})";
			//$nextOrder = 'sort ASC, id DESC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("((A.sort=%d AND A.id < %d) OR A.sort > %d)", nrs.Sort, nrs.ID, nrs.Sort)).
				OrderBy("sort ASC, id DESC").
				Limit(1)
		case 8:
			//$nextWhere = "status=1 AND cateid={$cateid} AND title LIKE '{$prefix}%' AND ((sort={$sort} AND id<{$id}) OR sort>{$sort})";
			//$nextOrder = 'sort ASC, id DESC';
			qb.Select("A.*,AC.name").
				From("jl_article A").
				LeftJoin("jl_article_cate AC").
				On("A.cateid=AC.id").
				Where(fmt.Sprintf("A.status=1")).
				And(fmt.Sprintf("A.cateid=%d", nrs.Cateid)).
				And(fmt.Sprintf("A.title LIKE \"%s%%\"", nrs.Prefix)).
				And(fmt.Sprintf("((A.sort=%d AND A.id < %d) OR A.sort > %d)", nrs.Sort, nrs.ID, nrs.Sort)).
				OrderBy("sort ASC, id DESC").
				Limit(1)
		}
	}
	sql := qb.String()
	Logger.Debugf("NextArticle sql: %v", sql)
	var adrs ArticleDetailsStruct
	err := o.Raw(sql).QueryRow(&adrs)
	if err == orm.ErrNoRows {
		return nil, nil
	} else if err != nil {
		Logger.Errorf("NextArticle query failed %v", err)
		return nil, err
	}
	Logger.Debugf("NextArticle return value: %v", adrs)
	return &adrs, nil
}
