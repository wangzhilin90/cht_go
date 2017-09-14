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
    string checkPhoneUse(1:CheckPhoneUseRequestStruct requestObj), //�����ֻ��Ų�ѯjl_user������鵽��¼����1001�����򷵻�1000
	
	//sql = 'select id from jl_user where hsid=$hsid'
    i32 getUserIdByhsid(1:GetUserIdByhsidRequestStruct requestObj), //����hsid��ѯjl_user���ȡ���û���ID������ID
   
	//sql = 'update jl_user set phone=$phone where id=$user_id'
    string updatePhone(1:UpdatePhoneRequestStruct requestObj), //����jl_user���޸ĸ��û����ֻ���
}
