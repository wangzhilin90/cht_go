//获取帮助中心文章列表
namespace php Common.HelpList
namespace go helplist

struct HelpListRequestStruct{
	1:i32 status,
	2:i32 cateid,
	3:string chengHuiTongTraceLog
}

struct HelpListResultStruct{
	1:string title,
	2:string content
}
//SELECT A.title,A.content FROM jl_article A LEFT JOIN jl_article_cate AC ON A.cateid=AC.id WHERE A.status={$status} AND A.cateid={$cateid} ORDER BY A.id ASC

struct HelpListResponseStrcut{
	1:list<HelpListResultStruct> HelpList
}

service HelpListThriftService{
	HelpListResponseStrcut getHelpList(1:HelpListRequestStruct requestObj)
}