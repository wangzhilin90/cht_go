namespace php User.UserBank
namespace go userbank

//详情
struct UserBankDetailsRequestStruct {
    1:i32 user_id,
    2:string chengHuiTongTraceLog,
}

struct UserBankDetailsStruct {
    1:i32 id,
    2:i32 user_id,//会员id
    3:string name,//银行开户名
    4:string account,//账号
    5:i32 bank,//开户银行名（id对应字典表）
    6:string branch,//银行支行名称
    7:i32 province,//省份id
    8:i32 city,//城市id
    9:i32 area,//区域id
    10:i32 addtime,//添加时间或者更新时间
    11:string addip,//添加时间或者更新ip
}

struct UserBankDetailsResponseStruct {
    1:i32 status,
    2:UserBankDetailsStruct UserBankDetails,
    3:string msg
}

//更新
struct UserBankUpdateRequestStruct {
    1:i32 id,
    2:i32 user_id,//会员id
    3:string name,//银行开户名
    4:string account,//账号
    5:i32 bank,//开户银行名（id对应字典表）
    6:string branch,//银行支行名称
    7:i32 province,//省份id
    8:i32 city,//城市id
    9:i32 area,//区域id
    10:i32 addtime,//添加时间或者更新时间
    11:string addip,//添加时间或者更新ip
    12:string chengHuiTongTraceLog,
}

struct UserBankUpdateResponseStruct {
    1:i32 status,
    2:string msg
}

//插入
struct UserBankInsertRequestStruct {
    1:i32 id,
    2:i32 user_id,//会员id
    3:string name,//银行开户名
    4:string account,//账号
    5:i32 bank,//开户银行名（id对应字典表）
    6:string branch,//银行支行名称
    7:i32 province,//省份id
    8:i32 city,//城市id
    9:i32 area,//区域id
    10:i32 addtime,//添加时间或者更新时间
    11:string addip,//添加时间或者更新ip
    12:string chengHuiTongTraceLog,
}

struct UserBankInsertResponseStruct {
    1:i32 status,
    2:string msg
}

//统计数量
struct UserBankCountRequestStruct {
    1:i32 user_id,
    2:string chengHuiTongTraceLog,
}

struct UserBankCountResponseStruct {
    1:i32 status,
    2:i32 total_num,
    3:string msg
}

service UserBankThriftService {
    UserBankDetailsResponseStruct getUserBankDetails (1:UserBankDetailsRequestStruct requestObj),
    UserBankUpdateResponseStruct updateUserBank  (1:UserBankUpdateRequestStruct requestObj),
    UserBankInsertResponseStruct insertUserBank  (1:UserBankInsertRequestStruct requestObj),
    UserBankCountResponseStruct  getUserBankNum  (1:UserBankCountRequestStruct requestObj),
}

//对应的sql:

//SELECT id,user_id,name,account,bank,branch,province,city,area .... FROM jl_user_bank WHERE user_id = $user_id LIMIT 1
//INSERT INTO jl_user_bank () VALUES ()
//UPDATE jl_user_bank SET .... WHERE id = 1
//SELECT COUNT(1) AS total_num FROM jl_user_bank UB LEFT JOIN jl_user U ON UB.user_id=U.id WHERE U.user_id = $user_id LIMIT 1