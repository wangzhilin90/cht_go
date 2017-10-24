//客户管理---商品管理---删除商品
namespace php Goods.GoodsDeL
namespace go goodsdel

struct GoodsDeLRequestStruct{
	1:i32 id,
	2:string chengHuiTongTraceLog
}

struct GoodsDeLResponseStruct{
	1:i32 status, //状态码
	2:string msg
}

service GoodsDeLThriftService {
	GoodsDeLResponseStruct delGoods (1:GoodsDeLRequestStruct requestObj)
}

//DELETE FROM `jl_point_goods` WHERE `id`='9' AND `sold_num`='0'