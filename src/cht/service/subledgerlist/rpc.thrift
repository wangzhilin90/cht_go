namespace go subledgerlist
namespace php User.SubledgerList

//请求结构体
struct SubledgerListRequestStruct {
	1:string hs_zhuanrangren_str ,//user_id,用逗号隔开，例如：160860,223082，233223
	2:string chengHuiTongTraceLog
}

//响应结构体
struct SubledgerDetailsStruct {
	1:i32 user_id, 	   //用户ID
	2:string realname, //真实姓名
	3:string card_id 	   //身份证号
}

struct SubledgerListResponseStruct {
	1:i32 status,
	2:string msg
	3:list<SubledgerDetailsStruct> SubledgerList
}

service SubledgerListThriftService {
    SubledgerListResponseStruct  getSubledgerList (1:SubledgerListRequestStruct requestObj)
}