namespace php Common.AdvertManage
namespace go advertmanagethriftservice

struct AdvertManageRequestStruct {
    1:i32 type,
    2:i32 limit,
    3:string chengHuiTongTraceLog
}

struct AdvertManageStruct {
    1:i32 id,
    2:i32 type, //广告位置：对应字典45类值；（原值：1注册；2登录；3首页右上角；4首页下方；5项目列表；6项目详情；7个人中心）
    3:string img,
    4:string adverturl,
    5:string title,
    6:i32 addtime,
    7:i32 adduser,
    8:i32 status,
    9:i32 fid,
    10:i32 starttime,
    11:i32 endtime,
}

struct AdvertManageResponseStruct {
	1:i32 status  //1000:查询广告成功 1001:查询广告失败
	2:string msg
    3:list<AdvertManageStruct> advertManageList
}

service AdvertManageThriftService {
    AdvertManageResponseStruct getAdvertManage(1:AdvertManageRequestStruct requestObj)
}

//查询 jl_advert_manage 表
//where条件：starttime <=  time() and endtime >= time() 
//type = type
//order by排序：addtime DESC
//如果limit=1取一条数据，否则按条件取所有数据
