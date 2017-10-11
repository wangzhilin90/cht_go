//运营数据
namespace php Common.OperationalData
namespace java net.cht.um.app.thrift.model.OperationalData
namespace go operationaldata

struct OperationalDataRequestStruct{
	1:i32 startMonth,
	2:i32 start,
	3:string chengHuiTongTraceLog
}

//最近30天投标排行
struct ThirtyDaysResultStruct{
	1:string money,
	2:string username
}
//SELECT SUM(BT.account_act) AS money,U.username FROM #@_borrow_tender BT LEFT JOIN #@_user U ON BT.user_id=U.id WHERE BT.status=1 AND BT.addtime>={strtotime('-30 day 00:00:00')} AND BT.addtime<={strtotime('23:59:59')} AND U.isvest=0 GROUP BY BT.user_id ORDER BY money DESC LIMIT 10

//最近12个月每月成交量
struct TwelveMonthResultStruct{
	1:string category,
	2:string account
}
//SELECT FROM_UNIXTIME(`review_time`,'%Y-%m') AS category,SUM(account) AS account FROM #@_borrow WHERE review_time>={$startMonth} AND status IN (2,3,6,7) GROUP BY category ORDER BY category

//最近1个月每月成交量
struct OneMonthResultStruct{
	1:string category,
	2:string account
}
//SELECT FROM_UNIXTIME(`review_time`,'%Y-%m') AS category,SUM(account) AS account FROM #@_borrow WHERE review_time>={$start} AND status IN (2,3,6,7) GROUP BY category ORDER BY category LIMIT 1

//借款周期占比
struct PeriodResultStruct{
	1:string category,
	2:string column_1
}
//SELECT time_limit AS category,COUNT(1) AS 'column_1' FROM #@_borrow WHERE status IN (2,3,6,7) AND addtime>1420041600 GROUP BY category ORDER BY category

//投资金额占比
struct InvestResultStruct{
	1:string a1,
	2:string a2,
	3:string a3,
	4:string a4,
	5:string a5
}
//SELECT SUM(IF(account_act<10000,`account_act`,0))/10000 AS a1, SUM(IF(account_act>=10000 AND account_act<100000,`account_act`,0))/10000 AS a2, SUM(IF(account_act>=100000 AND account_act<500000,`account_act`,0))/10000 AS a3, SUM(IF(account_act>=500000 AND account_act<1000000,`account_act`,0))/10000 AS a4, SUM(IF(account_act>=1000000,`account_act`,0))/10000 AS a5 FROM #@_borrow_tender WHERE status=1 LIMIT 1

//标的比例
struct BidResultStruct{
	1:i32 borrow_type,
	2:string number
}
//SELECT `borrow_type`,COUNT(1) AS number FROM #@_borrow WHERE status IN (2,3,6,7) GROUP BY borrow_type

//实时待收排行榜
struct WaitResultStruct{
	1:string money,
	2:string username
}
//SELECT U.username,A.hsmoney_wait AS money FROM #@_user U LEFT JOIN #@_account A ON A.user_id=U.id WHERE U.isvest=0 ORDER BY A.hsmoney_wait DESC LIMIT 10


struct OperationalDataResponseStruct {
	1:i32 status,
	2:string msg,
    3:list<ThirtyDaysResultStruct> ThirtyDaysList,
	4:list<TwelveMonthResultStruct> TwelveMonthList,
	5:list<OneMonthResultStruct> OneMonthList,
	6:string oldSum,//12个月之前成交总量 SELECT SUM(account) FROM #@_borrow WHERE review_time<{$startMonth} AND status IN (2,3,6,7) LIMIT 1
	7:string repayment,//目前累计成功还款 SELECT SUM(replayment_money) FROM #@_borrow_repayment WHERE status=1 LIMIT 1
	8:list<PeriodResultStruct> PeriodList,//借款周期占比
	9:InvestResultStruct InvestAccount,//投资金额占比
	10:list<BidResultStruct> BidList,//标的比例
	11:list<WaitResultStruct> WaitList//实时待收排行榜
}

service OperationalDataThriftService{
	OperationalDataResponseStruct getOperationalData(1:OperationalDataRequestStruct requestObj)
}
