namespace go gettenderredbagthriftservice

struct TenderRedbagRequestStruct {
    1:i32 userId, 
    2:i32 tenderId, //标ID，获取borrow_type
    3:i32 redId, //红包ID，查询红包的信息
    4:string tenderMoney,//用户投资金额，要求大于等于最小投资金额 ，小于等于最大金额
    5:i32  timeLimit, //投资期限
    6:string chengHuiTongTraceLog 
 }

struct TenderRedbagResponseStruct {
    1: i32 status, 
    2: string redbagMoney, //红包金额
    3: string msg
}

service GetTenderRedbagThriftService {
    TenderRedbagResponseStruct getRedbagInfo(1:TenderRedbagRequestStruct requestObj)
}