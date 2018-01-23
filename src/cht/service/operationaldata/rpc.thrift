//运营数据
namespace php Common.OperationalData
namespace go operationaldata

struct OperationalDataRequestStruct{
	1:i32 startMonth,
	2:i32 start,
	3:i32 today_time,
	4:i32 yesterday_time,
	5:i32 tomorrow_time,
	6:string chengHuiTongTraceLog
}

//最近30天投标排行
struct ThirtyDaysResultStruct{
	1:string money,
	2:string username
}

//最近12个月每月成交量
struct TwelveMonthResultStruct{
	1:string category,
	2:string account
}

//最近1个月每月成交量
struct OneMonthResultStruct{
	1:string category,
	2:string account
}

//借款周期占比
struct PeriodResultStruct{
	1:string category,
	2:string column_1
}

//出借金额占比
struct InvestResultStruct{
	1:string a1,
	2:string a2,
	3:string a3,
	4:string a4,
	5:string a5
}

//标的比例
struct BidResultStruct{
	1:i32 borrow_type,
	2:string number
}

//实时待收排行榜
struct WaitResultStruct{
	1:string money,
	2:string username
}

//投标统计
struct SumResultStruct{
    1:string tender,
    2:string tender_today,
    3:string profit
}


struct OperationalDataResponseStruct {
	1:i32 status,
	2:string msg,
    3:list<ThirtyDaysResultStruct> ThirtyDaysList,
	4:list<TwelveMonthResultStruct> TwelveMonthList,
	5:list<OneMonthResultStruct> OneMonthList,
	6:string oldSum,//12个月之前成交总量
	7:string repayment,//目前累计成功还款
	8:list<PeriodResultStruct> PeriodList,//借款周期占比
	9:InvestResultStruct InvestAccount,//出借金额占比
	10:list<BidResultStruct> BidList,//标的比例
	11:list<WaitResultStruct> WaitList,//实时待收排行榜
	12:SumResultStruct Sum,//投标统计
	13:i32 tenderUserCount//出借人数
}

service OperationalDataThriftService{
	OperationalDataResponseStruct getOperationalData(1:OperationalDataRequestStruct requestObj)
}
