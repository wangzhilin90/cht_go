namespace go  couponlistthriftservice
namespace php um.coupon
namespace java net.cht.um.app.thrift.model.regist

struct CouponRequestStruct {
    1: i32 user_id, 
	2: i16 status,
	3: i16 limit,
    4: string chengHuiTongTraceLog，
	5: string order_by  
}

struct CouponStruct {
    1: i32 id,
    2: i32 user_id,
    3: i32 addtime,
    4: i32 start_time,
    5: i32 end_time,
    6: i32 use_time,
    7: i16 status,
    8: i32 tender_id,
    9: string apr,
    10: string app_add,
    11: string min_tender,
    12: string max_tender,
    13: string time_limit,
    14: string borrow_type,
    15: string name,
    16: string remark,
    17: string activity_name
}

struct CouponResponseStruct {
     1: list<CouponStruct> couponList
}

service CouponListThriftService {
    CouponResponseStruct getCoupon(1:CouponRequestStruct requestObj)
}