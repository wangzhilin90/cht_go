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
    i32 checkEmailUse(1:CheckEmailUseRequestStruct requestObj), //where������email = $email and id != $user_id��ѯjl_user�������ѯ���û�˵�������ѱ�ʹ�÷���1�����򷵻�0

    i32 userAttestationSave(1:UserAttestationSaveStruct requestObj),
    // ���ܲ���
    // $data = array(
    //     'user_id' => $user_id,
    //     'email_status' => $email_status
    // );
    // if($email_status == 2) {
    //     $data['email_passtime'] = time();
    // }
    // ����user_id��ѯ��֤���Ƿ���ڸ��û�����֤���ݣ�������ھ͸������ݣ������������
    // ���»����ɹ�����1�����򷵻�0

    i32 userEmailSave(1:UserEmailSaveRequestStruct requestObj), //����user_id�޸��û���email���ɹ�����1��ʧ��0

    i32 sendEmail(1:SendEmailRequestStruct requestObj),
    // ���ܲ���
    // ����jl_sendmsg���������ݳɹ��õ�����ID
    // �����ʼ�
    // ���ͳɹ���������Ĳ���ID�޸�����status�ֶ�ֵΪ1��posttime�ֶ�Ϊϵͳ��ǰʱ�䣬����1
    // ����ʧ�ܷ���0

}