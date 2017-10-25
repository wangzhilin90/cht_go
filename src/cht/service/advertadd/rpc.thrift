//文章管理---广告图片管理---添加广告图片
namespace php Article.AdvertAdd
namespace go advertadd

struct AdvertAddRequestStruct{
	1:i32 type,
	2:string img,
	3:string adverturl,
	4:string title,
	5:i32 adduser,
	6:i32 fid,
	7:i32 starttime,
	8:i32 endtime,
	9:string chengHuiTongTraceLog
}

struct AdvertAddResponseStruct{
	1:i32 status,
	2:string msg
}

service AdvertAddThriftService {
	AdvertAddResponseStruct addAdvert (1:AdvertAddRequestStruct requestObj)
}

//addtime 获取当前系统时间