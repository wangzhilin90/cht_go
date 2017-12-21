//用户登录记录日志服务
namespace php Log.UpdateUserLoginLogDetails
namespace go updateuserloginlogdetails

struct UpdateUserLoginLogDetailsRequestStruct {
    1:i32 user_id,    //用户ID
    2:string login_ip,  //登录ip
    3:i32 login_style,  //登录方式0-PC，1-iOS，2-Android，3-WAP，4-微信'
    4:string chengHuiTongTraceLog,
}

struct UpdateUserLoginLogDetailsResponseStruct  {
     1:i32 status  //返回状态 1001 "更新登录日志失败"  1000 "更新登录日志成功" 
     2:string msg
}

struct UserLoginLogDetailsRequestStruct {
    1:i32 user_id,    //用户ID
    2:string login_style,  //(3,4) 或者 (0)
    3:string chengHuiTongTraceLog,
}

struct UserLoginLogDetailsStruct {
    1:i32 id,
    2:i32 user_id,
    3:i32 login_time,
    4:i32 login_style,
    5:string login_ip,
    6:string tender_money,
    7:i32 tender_time,
}

struct UserLoginLogDetailsResponseStruct {
    1:i32 status,
    2:string msg,
    3:UserLoginLogDetailsStruct userLoginLogDetails
}

service UpdateUserLoginLogDetailsThriftService {
    UpdateUserLoginLogDetailsResponseStruct updateUserLoginLogDetails (1:UpdateUserLoginLogDetailsRequestStruct requestObj),
    UserLoginLogDetailsResponseStruct getUserLoginLogDetails(1:UserLoginLogDetailsRequestStruct requestObj),//获取上次登录信息
}
//获取上次登录信息sql="select * from jl_user_login_log where user_id={$user_id} order by id desc limit 1"