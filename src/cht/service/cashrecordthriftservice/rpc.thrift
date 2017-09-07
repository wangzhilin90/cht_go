namespace go  cashrecordthriftservice

struct  CashRecordRequestStruct {
	1:i32 user_id 
	2:i32 start_time
	3:i32 end_time    
	4:i32 query_time //查询天数, 0:查全部  1：查最近7天 2：查一个月 3：查两个月
	5:i32 recharge_status //充值状态,0:查全部 1:已成功 2:审核中  3:审核失败
	6:i32 limit_offset  //查询偏移量
	7:i32 limit_num     //查询数量
	8:string chengHuiTongTraceLog
}

struct  CashRecordStruct{
	1:i32 id 
	2:i32 user_id
	3:string order_sn
	4:string money
	5:string credited
	6:string fee
	7:string use_return_money
	8:i32 use_free_num
	9:i32 addtime
	10:i32 status
	11:i32 pay_way
	12:i32 deal_time
	13:string  fail_result
}

struct CashRecordResponseStruct {
	1:i32 status  //1001:查询提现记录失败 1002 查询提现记录成功
	2:string msg
	3:i32 totalnum
	4:list<CashRecordStruct> cashRecordList	
}

service CashRecordThriftService {
	CashRecordResponseStruct getCashRecord(1:CashRecordRequestStruct requestObj)
}