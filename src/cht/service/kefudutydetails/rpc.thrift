//客户管理---客服值班---值班详情
namespace php User.KefuDutyDetails
namespace go  kefudutydetails

struct KefuDutyDetailsRequestStruct {
    1:i32 id,
	2:string chengHuiTongTraceLog
}

struct KefuDutyDetailsStruct {
	1:i32 id,
	2:string customer,
	3:i32 duty_time,
	4:string holiday_user,
	5:i32 is_rest,
	6:i32 starttime,//对应数据库"start"
	7:i32 endtime, //对应数据库"end"
}

struct KefuDutyDetailsResponseStruct{
	1:i32 status,
	2:string msg
    3:KefuDutyDetailsStruct KefuDutyDetails,
}

service KefuDutyDetailsThriftService {
    KefuDutyDetailsResponseStruct getKefuDutyDetails (1:KefuDutyDetailsRequestStruct requestObj)
}