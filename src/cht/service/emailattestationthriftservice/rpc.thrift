namespace php Common.EmailAttestation
namespace go emailattestationthriftservice

//邮箱认证服务

struct CheckEmailUseRequestStruct {
    1:string email,
    2:i32 user_id,
    3:string chengHuiTongTraceLog
}

struct UserAttestationSaveStruct {
    1:i32 user_id,
    2:i32 email_status,
    3:string chengHuiTongTraceLog
}

struct UserEmailSaveRequestStruct {
    1:string email,
    2:i32 user_id,
    3:string chengHuiTongTraceLog
}

struct SendEmailRequestStruct {
    1:i32 user_id,
    2:string send_to,
    3:string subject,
    4:string content,
    5:string ip,
    6:i32 addtime,
    7:string chengHuiTongTraceLog
}

service EmailAttestationThriftService {
    i32 checkEmailUse(1:CheckEmailUseRequestStruct requestObj), 
    i32 userAttestationSave(1:UserAttestationSaveStruct requestObj),
    i32 userEmailSave(1:UserEmailSaveRequestStruct requestObj), 
    i32 sendEmail(1:SendEmailRequestStruct requestObj),
}