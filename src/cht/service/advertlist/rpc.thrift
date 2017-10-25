//文章管理---广告图片管理---列表
namespace php Article.AdvertList
namespace go  advertlist

struct AdvertListRequestStruct{
	1:string chengHuiTongTraceLog
}

struct AdvertListStruct{
	1:i32 id,
	2:i32 type,
	3:string img,
	4:string adverturl,
	5:string title,
	6:i32 addtime,
	7:i32 adduser,
	8:i32 status,
	9:i32 fid,
	10:i32 starttime,
	11:i32 endtime
}

struct AdvertListResponseStruct{
	1:i32 status,
	2:string msg,
	3:list<AdvertListStruct> AdvertList
	4:i32 total_num,
}

service AdvertListThriftService {
	AdvertListResponseStruct getAdvertList (1:AdvertListRequestStruct requestObj)
}

//SELECT * FROM jl_advert_manage