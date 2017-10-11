namespace php Common.Talking
namespace go talking

//小诚交流日

struct TalkingRequestStruct{
	1:i32 cateid,
	2:i32 status,
	3:string chengHuiTongTraceLog
}

struct TalkListResultStruct{
	1:i32 id,
	2:string title,
	3:string img_url,
	4:string content
}

struct TalkingResponseStruct{
	1:list<TalkListResultStruct> TalkList,
	2:list<TalkListResultStruct> OneList
}

service TalkingThriftService{
	TalkingResponseStruct getTalkingList (1:TalkingRequestStruct requestObj)
}

//SELECT A.id,A.title,A.img_url,A.content FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.cateid=cateid AND A.status=status AND A.title LIKE '%交流会%' ORDER BY A.id DESC LIMIT 4
//SELECT A.id,A.title,A.img_url,A.content FROM #@_article A LEFT JOIN #@_article_cate AC ON A.cateid=AC.id WHERE A.id=3760 ORDER BY A.sort,A.addtime desc
