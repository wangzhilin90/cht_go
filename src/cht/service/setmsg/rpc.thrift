//用户设置提醒表
namespace php Common.SetMsg
namespace go  setmsg

//用户设置提醒详情请求结构体
struct SetMsgDetailsRequestStruct {
	1:i32 user_id,//用户id
	3:string chengHuiTongTraceLog
 }

struct SetMsgDetailsStruct {
	1:i32 id,
	2:i32 user_id,
	3:i32 addtime,
	4:i32 status,
	5:i32 tender_status,
	6:i32 borrow_status,
	7:i32 proto_status,
	8:i32 station_status,
	9:i32 guide_status,
	10:i32 sound_status
}

//用户设置提醒详情响应结构体
struct SetMsgDetailsResponseStruct {
	1:i32 status,
	2:SetMsgDetailsStruct SetMsgDetails,
	3:string msg
 }

//用户设置提醒处理的请求结构体
struct SetMsgDealRequestStruct {
	1:i32 user_id,
	2:i32 addtime,
	3:i32 status,
	4:i32 tender_status,
	5:i32 borrow_status,
	6:i32 proto_status,
	7:i32 station_status,
	8:i32 guide_status,
	9:i32 sound_status
	10:string chengHuiTongTraceLog
 }

 //用户设置提醒操作响应结构体
struct SetMsgDealResponseStruct {
	1:i32 status,
	2:string msg
 }

service SetMsgThriftService{
	SetMsgDetailsResponseStruct getSetMsgDetails(1:SetMsgDetailsRequestStruct requestObj) //基于user_id查询
	SetMsgDealResponseStruct   updateSetMsgDetails(1:SetMsgDealRequestStruct requestObj) //基于user_id 更新
	SetMsgDealResponseStruct   insertSetMsgDetails(1:SetMsgDealRequestStruct requestObj)
}

//详情：
//    SELECT 表中所有字段 FROM jl_set_msg WHERE user_id = $user_id ORDER BY id DESC LIMIT 1;
//    $checkCount = DB::selectOne('SELECT COUNT(1) AS num FROM jl_set_msg WHERE user_id = ?',array($userid))->num;//判断当前用户是否有记录
//    if ($checkCount) {
//    	$res = DB::update("UPDATE jl_set_msg SET {$field} = ?,addtime = ? WHERE user_id = ?",array($status,SYS_TIME,$userid));//更新记录
//    } else {
//    	$res = DB::insert("INSERT INTO jl_set_msg ({$field},user_id,addtime,id) VALUES (?,?,?,?)",array($status,$userid,SYS_TIME,"next VALUE FOR MYCATSEQ_SET_MSG"));
//    }
