namespace  go messagethriftservice 

//����ṹ��
struct MessageRequestStruct {
	1:i32 smsid ,//����id
	2:string phone ,//�ֻ���
	3:string addtime, //���ʱ��
	4:i32 type, //����     
	5:string chengHuiTongTraceLog
}

//��Ӧ�ṹ��
struct MessageInfoStruct {
	1:i32 id,
	2:i32 type,
	3:i32 user_id,
	4:string send_to ,
	5:string subject ,
	6:string content ,
	7:string attachment,
	8:i32 addtime,
	9:string ip,
	10:i32 posttime,
	11:i32 status
}

struct MessageInfoResponseStruct {
	1:i32 status,	//״̬ 1000���سɹ��� 1001����ʧ��
	2:string msg,	
	3:MessageInfoStruct MessageInfo //��ȡ��������
}

struct MessageCountResponseStruct {
	1:i32 status,	//״̬ 1000���سɹ��� 1002����ʧ��
	2:string msg,		
	3:i32 count ,	//��ȡ���ż�¼��
}

struct UserInfoStruct {
	1:i32 id,  //�û�ID
	2:string phone, //�ֻ���	
}

struct UserInfoResponseStruct {
	1:i32 status,//״̬ 1000���سɹ��� 1003����ʧ��
	2:string msg,
	3:UserInfoStruct userInfo,
}

service MessageThriftService {
    	MessageInfoResponseStruct getMessageInfo(1:MessageRequestStruct requestObj)
	MessageCountResponseStruct getMessageCount(1:MessageRequestStruct requestObj)
	UserInfoResponseStruct getUserInfo(1:MessageRequestStruct requestObj)
}