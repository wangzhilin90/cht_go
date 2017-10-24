//客户管理---商品管理---编辑商品
namespace php Goods.GoodsEdit
namespace go  goodsedit

struct GoodsEditRequestStruct{
	1:i32 id,
	2:i32 show_time,
	3:i32 close_time,
	4:i32 is_timer,
	5:string litpic,
	6:string name,
	7:i32 category,
	8:string redbag_money,
	9:i32 original_point,
	10:i32 current_point,
	11:i32 total_num,
	12:i32 single_num,
	13:string content,
	14:string chengHuiTongTraceLog
}

struct GoodsEditResponseStruct{
	1:i32 status, //状态码
	2:string msg
}

service GoodsEditThriftService {
	GoodsEditResponseStruct editGoods (1:GoodsEditRequestStruct requestObj)
}