//客户管理---商品管理---商品详情
namespace php Goods.GoodsDetails
namespace go  goodsdetails

struct GoodsDetailsRequestStruct{
	1:i32 id,
	2:string chengHuiTongTraceLog
}

struct GoodsDetailsStruct {
	1:i32 id,
	2:i32 addtime,
	3:i32 show_time,
	4:i32 close_time,
	5:i32 is_timer,
	6:i32 category,
	7:string redbag_money,
	8:i32 original_point,
	9:i32 current_point,
	10:i32 total_num,
	11:i32 sold_num,
	12:i32 single_num,
	13:string name,
	14:string litpic,
	15:string content
}

struct GoodsDetailsResponseStruct{
	1:i32 status, //状态码
    2:GoodsDetailsStruct GoodsDetails
	3:string msg
}

service GoodsDetailsThriftService {
	GoodsDetailsResponseStruct getGoodsDetails (1:GoodsDetailsRequestStruct requestObj)
}