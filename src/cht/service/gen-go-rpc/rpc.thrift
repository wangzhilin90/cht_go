namespace php Common.ArticleCate
namespace go  articlecate

//详情
struct ArticleCateListRequestStruct {
    1:i32 id,
    2:string name,//like查询
    3:string keywords,//like查询
    4:string description,//like查询
    5:i32 pid=-1,//父级id
    6:i32 status=-1,//是否显示，1:显示，0:不显示
    7:string chengHuiTongTraceLog,
}

struct ArticleCateDetailsStruct {
    1:i32 id,
    2:string name,//like查询
    3:string keywords,//like查询
    4:string description,//like查询
    5:i32 pid,//父级id
    6:i32 status,//是否显示，1:显示，0:不显示
    7:string img_url,
    8:i32 sort,
    9:i32 addtime
}

struct ArticleCateListResponseStruct {
    1:i32 status,
    2:list<ArticleCateDetailsStruct> ArticleCateList,
    3:string msg
}

service ArticleCateThriftService {
    ArticleCateListResponseStruct getArticleCateList (1:ArticleCateListRequestStruct requestObj),
}

//示意：SELECT id,name FROM jl_article_cate WHERE pid=14 ORDER BY id ASC

