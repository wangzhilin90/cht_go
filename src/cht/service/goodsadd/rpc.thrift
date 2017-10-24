//客户管理---商品管理---添加商品
namespace php Goods.GoodsAdd
namespace go goodsadd

struct GoodsAddRequestStruct{
	1:i32 show_time,
	2:i32 close_time,
	3:i32 is_timer,
	4:string litpic,
	5:string name,
	6:i32 category,
	7:string redbag_money,
	8:i32 original_point,
	9:i32 current_point,
	10:i32 total_num,
	11:i32 single_num,
	12:string content,
	13:string chengHuiTongTraceLog
}

struct GoodsAddResponseStruct{
	1:i32 status, //状态码
	2:string msg
}

service GoodsAddThriftService {
	GoodsAddResponseStruct addGoods (1:GoodsAddRequestStruct requestObj)
}