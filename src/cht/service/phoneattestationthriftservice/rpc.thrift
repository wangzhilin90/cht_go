namespace php Common.PhoneAttestation
namespace go phoneattestationthriftservice

//手机认证服务

struct CheckPhoneUseRequestStruct {
    1:string phone,
    2:string chengHuiTongTraceLog
}
struct GetUserIdByhsidRequestStruct {
    1:string hsid,
    2:string chengHuiTongTraceLog
}

struct UpdatePhoneRequestStruct {
    1:string phone,
    2:i32 user_id,
    3:string chengHuiTongTraceLog
}

service PhoneAttestationThriftService {
    string checkPhoneUse(1:CheckPhoneUseRequestStruct requestObj), //根据手机号查询jl_user表，如果查到记录返回1001，否则返回1000
    i32 getUserIdByhsid(1:GetUserIdByhsidRequestStruct requestObj), //根据hsid查询jl_user表获取该用户的ID，返回ID
    string updatePhone(1:UpdatePhoneRequestStruct requestObj), //更新jl_user表修改该用户的手机号, 1000:更新手机号成功 1001:查找原手机号失败 1002：新手机号与原手机号一致 1003:插入日志失败 1004:更新手机号失败
}