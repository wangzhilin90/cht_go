namespace go gettendercouponthriftservice

struct TenderCouponRequestStruct  {
    1:i32 userId, 
    2:i32 tenderId, //标ID，获取borrow_type
    3:i32 couponId, //加息券ID，查询加息券的信息
    4:string tenderMoney,//用户投资金额，要求大于等于最小投资金额 ，小于等于最大金额
    5:i32  timeLimit, //投资期限
    6:string chengHuiTongTraceLog 
 }

struct TenderCouponResponseStruct  {
    1: i32 status, 
    2: string coupon, //加息值，比如前台显示0.12%，这里存0.0012'
    3: string msg
}

service GetTenderCouponThriftService {
    TenderCouponResponseStruct  getCouponInfo(1:TenderCouponRequestStruct  requestObj)
}