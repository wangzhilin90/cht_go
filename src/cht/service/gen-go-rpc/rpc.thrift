namespace go loguserloginservice

struct LogUserlLoginRequestStruct {
    1:i32 user_id,    //用户ID
	2:string login_ip,	//登录ip
	3:i32 login_style,  //登录方式0-PC，1-iOS，2-Android，3-WAP，4-微信'
    4:string chengHuiTongTraceLog,
}

struct LogUserLoginResponseStruct {
     1:i32 status  //返回状态 1001 "更新登录日志失败" 1002"更新登录日志成功" 
     2:string msg
}

service LogUserLoginThriftService {
    LogUserLoginResponseStruct updateLogUserlLogin(1:LogUserlLoginRequestStruct requestObj)
}
