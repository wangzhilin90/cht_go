namespace php User.UserAppBank
namespace go  userappbank

//app会员银行信息表详情
struct UserAppBankDetailsRequestStruct {
    1:i32 user_id,
    2:string chengHuiTongTraceLog,
}

struct UserAppBankDetailsStruct {
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

struct UserAppBankDetailsResponseStruct {
    1:i32 status,
    2:UserAppBankDetailsStruct UserAppBankDetails,
    3:string msg
}

//更新
struct UserAppBankUpdateRequestStruct {
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

struct UserAppBankUpdateResponseStruct {
    1:i32 status,
    2:string msg
}

//插入
struct UserAppBankInsertRequestStruct {
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

struct UserAppBankInsertResponseStruct {
    1:i32 status,
    2:string msg
}

//删除
struct UserAppBankDeleteRequestStruct {
    1:i32 user_id,//会员id
    2:string chengHuiTongTraceLog,
}

struct UserAppBankDeleteResponseStruct {
    1:i32 status,
    2:string msg
}

service UserAppBankThriftService {
    UserAppBankDetailsResponseStruct getUserAppBankDetails (1:UserAppBankDetailsRequestStruct requestObj),
    UserAppBankUpdateResponseStruct updateUserAppBank  (1:UserAppBankUpdateRequestStruct requestObj),
    UserAppBankInsertResponseStruct insertUserAppBank  (1:UserAppBankInsertRequestStruct requestObj),
    UserAppBankDeleteResponseStruct deletetUserAppBank (1:UserAppBankDeleteRequestStruct requestObj),
}


//select * from jl_user_appbank where user_id = 106075 limit 1;
//update jl_user_appbank set name = '134',account='' where user_id = $user_id
//insert into jl_user_appbank () values () ;
//delete * from jl_user_appbank where user_id = 234234