//文章详情
namespace php Common.ArticleDetails
namespace go articledetails

//文章内容请求结构体
struct ArticleDetailsRequestStruct{
	1:i32 id,
	2:i32 status=-1,//新增字段查询功能
	3:string chengHuiTongTraceLog

}

//文章内容返回结构体
struct ArticleDetailsResultStruct{
	1:i32 id,
	2:i32 cateid,
	3:string title,
	4:string content,
	5:string keywords,
	6:string description,
	7:string img_url,
	8:string sort,
	9:i32 status,
	10:i32 addtime,
	11:string banner_url,
	12:i32 isbanner,
	13:i32 type,
	14:string name,
	15:i32 result_status,
	16:string msg
}
//SELECT A.*,AC.name FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id WHERE A.id='3777'

//上、下一篇请求结构体
struct NextRequestStruct{
	1:i32 id,
	2:i32 cateid,
	3:i32 type,
	4:i32 addtime,
	5:string chengHuiTongTraceLog
}


service ArticleDetailsThriftService{
	ArticleDetailsResultStruct getArticleDetails(1:ArticleDetailsRequestStruct requestObj)
	i32 updateReadNum(1:ArticleDetailsRequestStruct requestObj)              //更新阅读量 UPDATE `jl_article` SET `read_num` = `read_num`+'1' WHERE id=3777
	ArticleDetailsResultStruct prevArticle(1:NextRequestStruct requestObj)  //上一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.id<3195 AND type=1 ORDER BY A.id desc
	ArticleDetailsResultStruct nextArticle(1:NextRequestStruct requestObj) //下一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.id>3195 AND type=1
}
//if (cateid == 10) //上一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.addtime<{addtime} ORDER BY A.addtime desc
//                  //下一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.addtime>{addtime}

