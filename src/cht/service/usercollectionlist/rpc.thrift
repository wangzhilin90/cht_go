//查询汇款明细服务
namespace php Finance.UserCollectionList
namespace go usercollectionlist

struct UserCollectionListRequestStruct {
	1:i32 user_id,//用户id
	2:i32 starttime,//开始时间
	3:i32 endtime,  //结束时间
	4:i32 search_time, //1:近7天，2:1个月，3:2个月
	5:i32 state, //0:全部，1: 还款中2：已回款
	6:i32 limitOffset,//偏移量
	7:i32 limitNum,//每页现实多长时间
	8:string borrowid,//项目编号
	9:i32 check_zhuanrangren,//检查转让人
	10:string chengHuiTongTraceLog
}

struct UserCollectionDetailsStruct {
	1:string username,  // U.username
	2:string title,		//B.title
	3:i32 is_datetype,   //B.is_datetype
	4:i32 time_limit,    //B.time_limit
	5:string zhuanrangren,  //B.zhuanrangren
	6:i32 repay_time,	//BC.repay_time
	7:i32 borrow_id,	//BT.borrow_id
	8:i32 periods,		//BC.periods
	9:i32 repay_yestime,	//BC.repay_yestime
	10:string repay_yesaccount,	//BC.repay_yesaccount
	11:string repay_account,	//BC.repay_account
	12:string capital,		//BC.capital
	13:string interest,		//BC.interest
	14:string late_interest,	//BC.late_interest
	15:i32 late_days,		//BC.late_days
	16:i32 status,			//BC.status
	17:string interest_add,	//BC.interest_add
	18:i32 old_user_id,		//BC.old_user_id
	19:i32 style			//B.style
}

struct UserCollectionListResponseStruct {
		1:i32 status,
		2:string msg,
		3:i32 totalNum,
		4:list<UserCollectionDetailsStruct> UserCollectionList,
}

service UserCollectionListThriftService {
    UserCollectionListResponseStruct getUserCollectionList (1:UserCollectionListRequestStruct requestObj)
}