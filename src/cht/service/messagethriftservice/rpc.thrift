namespace  go messagethriftservice 

//请求结构体
struct MessageRequestStruct {
	1:i32 smsid ,//短信id
	2:string phone ,//手机号
	3:string addtime, //添加时间
	4:i32 type, //类型     
	5:string chengHuiTongTraceLog
}

//响应结构体
struct MessageInfoStruct {
	1:i32 id,
	2:i32 type,
	3:i32 user_id,
	4:string send_to ,
	5:string subject ,
	6:string content ,
	7:string attachment,
	8:i32 addtime,
	9:string ip,
	10:i32 posttime,
	11:i32 status
}

struct MessageInfoResponseStruct {
	1:i32 status,	//状态 1000返回成功， 1001返回失败
	2:string msg,	
	3:MessageInfoStruct MessageInfo //获取短信详情
}

struct MessageCountResponseStruct {
	1:i32 status,	//状态 1000返回成功， 1002返回失败
	2:string msg,		
	3:i32 count ,	//获取短信记录数
}

struct UserInfoStruct {
	1:i32 id,  //用户ID
	2:string phone, //手机号	
}

struct UserInfoResponseStruct {
	1:i32 status,//状态 1000返回成功， 1003返回失败
	2:string msg,
	3:UserInfoStruct userInfo,
}

service MessageThriftService {
    MessageInfoResponseStruct getMessageDetails(1:MessageRequestStruct requestObj)
	MessageCountResponseStruct getMessageCount(1:MessageRequestStruct requestObj)
	UserInfoResponseStruct getUserInfo(1:MessageRequestStruct requestObj)
}