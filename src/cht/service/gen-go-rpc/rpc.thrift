namespace go  rechargerecordthriftservice

struct  RechargeRecordRequestStruct {
	1:i32 user_id 
	2:i32 start_time
	3:i32 end_time    
	4:i32 query_time //��ѯ����, 0:��ȫ��  1�������7�� 2����һ���� 3����������
	5:i32 recharge_status //��ֵ״̬,0:��ȫ�� 1:�ѳɹ� 2:�����  3:���ʧ��
	6:i32 limit_offset  //��ѯƫ����
	7:i32 limit_num     //��ѯ����
	8:string chengHuiTongTraceLog
}

struct  RechargeRecordStruct{
	1:i32 id 
	2:i32 user_id
	3:string order_sn
	4:string money
	5:i32 addtime
	6:i32 status
	7:i32 deal_time
	8:i32 pay_type
	9:i32 pay_way
	10:string  fail_result
}

struct RechargeRecordResponseStruct {
	1:i32 status  //0:��ѯ��ֵ��¼�ɹ� 1001 ��ѯ��ֵ��¼ʧ��
	2:string Msg
	3:i32 totalnum //��ֵ�ܼ�¼��
	4:string totalHsRechargeMoney //��ֵ�ܽ�� SELECT SUM(money) FROM jl_hs_recharge WHERE user_id = $user_id  AND status = 1
	5:list<RechargeRecordStruct> rechargeRecordList	
}

service RechargeRecordThriftService {
	RechargeRecordResponseStruct getRechargeRecord(1:RechargeRecordRequestStruct requestObj)
}
