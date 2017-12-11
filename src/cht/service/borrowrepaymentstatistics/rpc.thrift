namespace php Borrow.BorrowRepaymentStatistics
namespace go  borrowrepaymentstatistics

struct RepaymentStatisticsRequestStruct {
    1:i32 user_id,
    2:string chengHuiTongTraceLog
}

struct RepaymentStatisticsDetailsStruct {
    1:i32 borrow_id,
    2:string will_money,
    3:string replayment_money,
    4:string noreplayment_money,
}

struct RepaymentStatisticsResponseStruct {
    1:i32 status ,
    2:list<RepaymentStatisticsDetailsStruct> RepaymentStatisticsList,
    3:string msg
}

service BorrowRepaymentStatisticsThriftService {
    RepaymentStatisticsResponseStruct getRepaymentStatisticsDetails (1:RepaymentStatisticsRequestStruct requestObj),
}

//sql 及部分逻辑
//汇总尝还表 偿还本息 已还本息 	未还本息
//$repay = DB::select('SELECT borrow_id,SUM(will_money) AS will_money,SUM(replayment_money) AS replayment_money,SUM(will_money-replayment_money) AS noreplayment_money FROM jl_borrow_repayment WHERE user_id=? ORDER BY borrow_id GROUP BY borrow_id',[$userid]);