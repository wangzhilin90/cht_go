//文章管理---广告图片管理---修改广告图片
namespace php Article.AdvertUpdate
namespace go  advertupdate

struct AdvertUpdateRequestStruct{
	1:i32 id,
	2:i32 type,
	3:string img,
	4:string adverturl,
	5:string title,
	6:i32 adduser,
	7:i32 fid,
	8:i32 starttime,
	9:i32 endtime,
	10:string chengHuiTongTraceLog
}

struct AdvertUpdateResponseStruct{
	1:i32 status,
	2:string msg
}

service AdvertUpdateThriftService {
	AdvertUpdateResponseStruct updateAdvert (1:AdvertUpdateRequestStruct requestObj)
}

//UPDATE `jl_advert_manage` SET `img` = 'advert/2017/1019/20171019045638703.png',`adverturl` = '',`fid` = '59769',`type` = '2',`title` = '登录11',`starttime` = '1497801599',`endtime` = '1548950399',`adduser` = '1',`addtime` = '1508403399' WHERE `id`='1'
//addtime获取系统当前时间