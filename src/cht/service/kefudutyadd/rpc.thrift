//客户管理---客服值班---新增值班
namespace php User.KefuDutyAdd
namespace go  kefudutyadd

struct KefuDutyAddRequestStruct{
	1:string customer,
	2:i32 duty_time,
	3:string holiday_user,
	4:i32 is_rest,
	5:i32 addtime,
	6:i32 starttime, //对应数据库"start"
	7:i32 endtime, //对应数据库"end"
	8:string chengHuiTongTraceLog
}

struct KefuDutyAddResponseStruct {
	1:i32 status,
	2:string msg
}

service KefuDutyAddThriftService {
    KefuDutyAddResponseStruct addKefuDuty (1:KefuDutyAddRequestStruct requestObj)
}