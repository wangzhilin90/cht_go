namespace php User.SecuredList
namespace go securedlist

//做标---担保人服务
struct SecuredListRequestStruct {
	1:string chengHuiTongTraceLog
}

struct SecuredDetailsStruct {
	1:string secured
}

struct SecuredListResponseStruct {
	1:i32 status,
	2:string msg,
	3:list<SecuredDetailsStruct> SecuredList
}

service SecuredListThriftService {
    SecuredListResponseStruct getSecuredList (1: SecuredListRequestStruct requestObj)
}