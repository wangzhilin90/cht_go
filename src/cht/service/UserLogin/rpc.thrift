namespace go UserLogin

struct UserlLoginRequestStruct {
    1:string username, //用户名
    2:string password, //加密后的password
    3:string ip,       //用户当前登录ip
	4:i32  isadmin,   //是否为后台，默认是0前台登录
    5:string chengHuiTongTraceLog,
}

struct UserLoginResponseStruct {
    1:i32 user_id,  //用户名
    2:i32 status,  //status与msg对应  1001-密码重试次数太多；1002-帐号不存在，请重新输入！；1003-您的帐号被锁定了，请联系我们。；1004: "密码错误",1005: "密码验证通过",
    3:string msg,
    4:i32 flag,  //剩余登录次数
}

service UserLoginThriftService {
    UserLoginResponseStruct getUserLoginInfo(1:UserlLoginRequestStruct requestObj)
}