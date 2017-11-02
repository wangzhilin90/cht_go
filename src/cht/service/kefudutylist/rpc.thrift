//客户管理---客服值班---列表
namespace php User.KefuDutyList
namespace go  kefudutylist

struct KefuDutyListRequestStruct {
	1:i32 start_time,
	2:i32 end_time,
	3:i32 kefu,
	4:i32 is_export,//是否导出，默认不导出 0
	5:i32 limitOffset,
	6:i32 limitNum,
	7:string chengHuiTongTraceLog
}

struct KefuDutyListResultStruct {
	1:i32 id,
	2:string customer,
	3:i32 is_rest,
	4:i32 duty_time,
	5:string holiday_user,
	6:i32 addtime,
	7:i32 starttime, //对应数据库"start"
	8:i32 endtime //对应数据库"end"
}

struct KefuDutyListResponseStruct {
	1:i32 status,
	2:string msg,
	3:list<KefuDutyListResultStruct> KefuDutyList,
	4:i32 total_num
}

service KefuDutyListThriftService {
	KefuDutyListResponseStruct getKefuDutyList (1:KefuDutyListRequestStruct requestObj)
}

//总数：SELECT COUNT(1) FROM jl_customer_plan
//列表：SELECT * FROM jl_customer_plan LIMIT 0,15
//导出不加分页

/**
 * 客服值班列表
 */
//    if (isset($_GET['start_time']) && $_GET['start_time']) {
//         $start_time = strtotime($_GET['start_time']);
//         $where[] = "duty_time >= '$start_time'";
//     }
//     if (isset($_GET['end_time']) && $_GET['end_time']) {
//         $end_time = strtotime($_GET['end_time']);
//         if ($start_time <= $end_time) {
//             $end_time = strtotime("next day", $end_time);
//             $where[] = "duty_time < '$end_time'";
//         }
//     }
//     if (isset($_GET['kefu']) && $_GET['kefu']) {
//         $where[] = "customer = .$_GET['kefu']."%'";
//     }



