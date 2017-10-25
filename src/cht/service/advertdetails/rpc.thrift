//文章管理---广告图片管理---图片详情
namespace php Article.AdvertDetails
namespace go advertdetails

struct AdvertDetailsRequestStruct{
	1:i32 id,
	2:string chengHuiTongTraceLog
}

struct AdvertDetailsStruct{
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

struct AdvertDetailsReponseStruct{
	1:i32 status,
	2:string msg 
    3:AdvertDetailsStruct AdvertDetails
}

service AdvertDetailsThriftService {
	AdvertDetailsReponseStruct getAdvertDetails (1:AdvertDetailsRequestStruct requestObj)
}

//SELECT * FROM jl_advert_manage WHERE `id`='1' LIMIT 1