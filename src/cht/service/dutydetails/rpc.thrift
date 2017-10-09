namespace php SysUser.DutyDetails
namespace go dutydetails

struct DutyDetailsRequestStruct {
    1:string chengHuiTongTraceLog
}

//值班详情
struct DutyDetailsStruct {
  1:i32 id,
  2:string customer, //值班客服,
  3:i32 is_rest,    //是否为节假日，1：是；0：不是,
  4:i32 duty_time, //值班时间,
  5:string holiday_user,//休假客服,
  6:i32 start_time,//值班开始时间,(数据库名字对应start)
  7:i32 end_time,//值班结束时间(数据库名字对应end,直接用end不能生成代码)
  8:i32 addtime // 添加时间,
}

struct DutyDetailsResponseStruct {
    1:i32 status,
    2:DutyDetailsStruct DutyDetails,
    3:string msg
}

service DutyDetailsThriftService {
	// sql = "select * from jl_customer_plan where duty_time >= strtotime('00:00:00') and duty_time <= strtotime("23:59:59") and start <= time() and end >= time() limit 1"
    DutyDetailsResponseStruct getDutyDetails (1:DutyDetailsRequestStruct requestObj)
}
