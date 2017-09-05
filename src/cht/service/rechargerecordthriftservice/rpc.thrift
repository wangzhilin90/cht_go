namespace go  rechargerecordthriftservice

struct  RechargeRecordRequestStruct {
	1:i32 user_id 
	2:i32 start_time
	3:i32 end_time    
	4:i32 query_time //查询天数, 0:查全部  1：查最近7天 2：查一个月 3：查两个月
	5:i32 recharge_status=3 //充值状态,0:查全部 1:已成功 2:审核中  3:审核失败
	6:i32 limit_offset  //查询偏移量
	7:i32 limit_num     //查询数量
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
	1:i32 status  //0:查询充值记录成功 1001 查询充值记录失败
	2:string Msg
	3:i32 totalnum
	4:list<RechargeRecordStruct> rechargeRecordList	
}

service RechargeRecordThriftService {
	RechargeRecordResponseStruct getRechargeRecord(1:RechargeRecordRequestStruct requestObj)
}