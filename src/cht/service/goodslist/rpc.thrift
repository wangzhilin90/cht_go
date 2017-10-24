//客户管理---商品管理列表
namespace php Goods.GoodsList
namespace go  goodslist

struct GoodsListRequestStruct{
	1:string name,
	2:i32 category,
	3:i32 is_export,//是否导出，默认不导出 0
	4:i32 limitOffset,
	5:i32 limitNum,
	6:string chengHuiTongTraceLog
}

struct GoodsListResultStruct{
	1:i32 id,
	2:i32 addtime,
	3:i32 show_time,
	4:i32 close_time,
	5:i32 is_timer,
	6:i32 category,
	7:string redbag_money, //红包金额
	8:i32 original_point, //原始所需积分
	9:i32 current_point, //当前所需积分
	10:i32 total_num, //总数量
	11:i32 sold_num, //已售出数量
	12:i32 single_num, //单人最多可购买数量，0为不限
	13:string name, //商品名称
	14:string litpic, //商品图片
	15:string content //商品详情
}

struct GoodsListReponseStruct{
	1:i32 status,
	2:string msg,
	3:list<GoodsListResultStruct> GoodsList
	4:i32 total_num,
}

service GoodsListThriftService {
    GoodsListReponseStruct getGoodsList (1:GoodsListRequestStruct requestObj)
}

//		  $where = [];
//        if (!empty($_GET['name'])) {
//            $where[] = "name LIKE '%{$_GET['name']}%'";
//        }
//        if (!empty($_GET['category'])) {
//            $where['category'] = $_GET['category'];
//        }

//导出：SELECT * FROM jl_point_goods ORDER BY id DESC

//总数：SELECT COUNT(1) FROM jl_point_goods
//列表：SELECT * FROM jl_point_goods ORDER BY id DESC LIMIT 0,20








