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
struct ArticleDetailsStruct {
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
}

//文章内容返回结构体
struct ArticleDetailsResponseStruct {
    1:i32 status,
    2:ArticleDetailsStruct ArticleDetails,
    3:string msg
}
//SELECT A.*,AC.name FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id WHERE A.id='3777'

//上、下一篇请求结构体
struct NextRequestStruct{
	1:i32 id,
	2:i32 cateid,
	3:i32 type,
	4:i32 addtime,
	5:i32 sort,
	6:string prefix,
	7:i32 is_app = 0,       //请求上、下页接口是先判断此参数,如果is_app=1则为app端请求，请按下述APP文章上、下页逻辑判断
	8:string chengHuiTongTraceLog
}


service ArticleDetailsThriftService{
	ArticleDetailsResponseStruct getArticleDetails(1:ArticleDetailsRequestStruct requestObj)
	i32 updateReadNum (1:ArticleDetailsRequestStruct requestObj)              //更新阅读量 UPDATE `jl_article` SET `read_num` = `read_num`+'1' WHERE id=3777
	ArticleDetailsResponseStruct prevArticle(1:NextRequestStruct requestObj)  //上一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.id<3195 AND type=1 ORDER BY A.id desc
	ArticleDetailsResponseStruct nextArticle(1:NextRequestStruct requestObj) //下一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.id>3195 AND type=1
}
//if (cateid == 10) //上一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 AND A.status=1 AND A.addtime<{addtime} ORDER BY A.addtime desc
//                  //下一篇 SELECT A.*,AC.name FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=5 and A.status=1 AND A.addtime>{addtime}


//APP文章上、下页逻辑判断
//        if ($cateid == 4) {
//            $prevWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND addtime>{$addtime}) OR sort<{$sort})";
//            $prevOrder = 'sort DESC, addtime ASC';
//            $nextWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND addtime<{$addtime}) OR sort>{$sort})";
//            $nextOrder = 'sort ASC, addtime DESC';
//        } elseif ($cateid == 5) {
//            $prevWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND id>{$id}) OR sort<{$sort})";
//            $prevOrder = 'sort DESC, id ASC';
//            $nextWhere = "status=1 AND cateid={$cateid} AND ((sort={$sort} AND id<{$id}) OR sort>{$sort})";
//            $nextOrder = 'sort ASC, id DESC';
//        } elseif ($cateid == 8) {
//            $prevWhere = "status=1 AND cateid={$cateid} AND title LIKE '{$prefix}%' AND ((sort={$sort} AND id>{$id}) OR sort<{$sort})";
//            $prevOrder = 'sort DESC, id ASC';
//            $nextWhere = "status=1 AND cateid={$cateid} AND title LIKE '{$prefix}%' AND ((sort={$sort} AND id<{$id}) OR sort>{$sort})";
//            $nextOrder = 'sort ASC, id DESC';
//        }

