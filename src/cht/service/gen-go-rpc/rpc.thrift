namespace go phoneattestationthriftservice

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
    //sql = 'select count(1) from jl_user where phone=$phone'
    string checkPhoneUse(1:CheckPhoneUseRequestStruct requestObj), //根据手机号查询jl_user表，如果查到记录返回1001，否则返回1000
	
	//sql = 'select id from jl_user where hsid=$hsid'
    i32 getUserIdByhsid(1:GetUserIdByhsidRequestStruct requestObj), //根据hsid查询jl_user表获取该用户的ID，返回ID
   
	//sql = 'update jl_user set phone=$phone where id=$user_id'
    string updatePhone(1:UpdatePhoneRequestStruct requestObj), //更新jl_user表修改该用户的手机号
}
