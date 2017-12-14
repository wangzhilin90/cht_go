namespace php User.UserTimes
namespace go  usertimes

//用户登陆次数详情
struct UserTimesDetailsRequestStruct {
    1:string username,
    2:i32 isadmin,
	3:string chengHuiTongTraceLog
}

struct UserTimesDetailsStruct {
    1:string username,
    2:string ip,
    3:i32 logintime,
    4:i32 times,
    5:i32 isadmin
}

struct UserTimesDetailsResponseStruct {
    1:i32 status,
    2:UserTimesDetailsStruct UserTimesDetails
    3:string msg
}

//用户登陆次数更新
struct UserTimesUpdateRequestStruct {
    1:string username,
    2:string ip,
    3:i32 logintime,
    4:i32 times,
    5:i32 isadmin
	6:string chengHuiTongTraceLog
}

struct UserTimesUpdateResponseStruct {
    1:i32 status,
    2:string msg
}

//用户登陆次数新增
struct UserTimesInsertRequestStruct {
    1:string username,
    2:string ip,
    3:i32 logintime,
    4:i32 times,
    5:i32 isadmin
	6:string chengHuiTongTraceLog
}

struct UserTimesInsertResponseStruct {
    1:i32 status,
    2:string msg
}

service UserTimesThriftService {
	UserTimesDetailsResponseStruct getUserTimesDetails (1:UserTimesDetailsRequestStruct requestObj)//基于username获取详情
	UserTimesUpdateResponseStruct  updateUserTimes     (1:UserTimesUpdateRequestStruct requestObj)//基于username更新
	UserTimesInsertResponseStruct  insertUserTimes     (1:UserTimesInsertRequestStruct requestObj)
}
