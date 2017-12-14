namespace  php Common.Message
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
struct MessageDetailsStruct {
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

struct MessageDetailsResponseStruct {
	1:i32 status,	//状态 1000返回成功， 1001返回失败
	2:string msg,
	3:MessageDetailsStruct MessageDetails //获取短信详情
}

struct MessageCountResponseStruct {
	1:i32 status,	//状态 1000返回成功， 1002返回失败
	2:string msg,
	3:i32 count ,	//获取短信记录数
}

struct UserDetailsStruct {
	1:i32 id,  //用户ID
	2:string phone, //手机号
}

struct UserDetailsResponseStruct {
	1:i32 status,//状态 1000返回成功， 1003返回失败
	2:string msg,
	3:UserDetailsStruct userDetails,
}

//更新站内信信息
struct MessageUpdateRequestStruct {
	1:i32 to_user ,//接收人id  where条件1
	2:i32 is_push_flag_old=-1,//推送标识旧 where条件2
	3:i32 is_push_flag_new=-1 ,//推送标识新,更新字段
	4:string chengHuiTongTraceLog
}

//更新站内信信息
struct MessageUpdateResponseStruct {
    1:i32 status,
    2:string msg
}

service MessageThriftService {
    MessageDetailsResponseStruct getMessageDetails(1:MessageRequestStruct requestObj)
	MessageCountResponseStruct   getMessageCount  (1:MessageRequestStruct requestObj)
	UserDetailsResponseStruct    getUserDetials   (1:MessageRequestStruct requestObj)
	MessageUpdateResponseStruct    updateMessage    (1:MessageUpdateRequestStruct requestObj)
}

//update jl_message set is_push_flag = $is_push_flag_new where is_push_flag = $is_push_flag_old and to_user = $user_id;



