namespace go emailattestationthriftservice

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
    i32 checkEmailUse(1:CheckEmailUseRequestStruct requestObj), //where条件：email = $email and id != $user_id查询jl_user表如果查询到用户说明邮箱已被使用返回1，否则返回0

    i32 userAttestationSave(1:UserAttestationSaveStruct requestObj),
    // 接受参数
    // $data = array(
    //     'user_id' => $user_id,
    //     'email_status' => $email_status
    // );
    // if($email_status == 2) {
    //     $data['email_passtime'] = time();
    // }
    // 根据user_id查询认证表看是否存在该用户的认证数据，如果存在就更新数据，否则插入数据
    // 更新或插入成功返回1，否则返回0

    i32 userEmailSave(1:UserEmailSaveRequestStruct requestObj), //根据user_id修改用户的email，成功返回1，失败0

    i32 sendEmail(1:SendEmailRequestStruct requestObj),
    // 接受参数
    // 插入jl_sendmsg表，插入数据成功得到插入ID
    // 发送邮件
    // 发送成功根据上面的插入ID修改数据status字段值为1，posttime字段为系统当前时间，返回1
    // 发送失败返回0

}