//第三方支付方式列表
namespace php Common.PaymentConfigList
namespace go  paymentconfiglist

struct PaymentConfigListRequestStruct {
    1:string chengHuiTongTraceLog
}

struct PaymentConfigDetailsStruct {
    1:i32 id,//支付方式表,
    2:i32 type,//分类,
    3:string nid,//标识,
    4:string name,//支付方式名称,
    5:string logo,//COMMENT logo,
    6:string config,//配置信息,
    7:string fee,//手续费,
    8:i32 status,//状态 0开通 1停用,
    9:string remark,//备注,
    10:i32 sort,//排序,
}

struct PaymentConfigListResponseStruct {
    1:i32 status,
    2:string msg,
    3:list<PaymentConfigDetailsStruct> PaymentConfigList
}

service PaymentConfigListThriftService {
    PaymentConfigListResponseStruct getPaymentConfigList (1:PaymentConfigListRequestStruct requestObj)
}

//sql = "select * from jl_payment_config order by sort"