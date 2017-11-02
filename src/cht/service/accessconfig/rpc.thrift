//推广名称记录表
namespace php Common.AccessConfig
namespace go  accessconfig

struct AccessConfigRequestStruct{
	1:string source,
	2:string chengHuiTongTraceLog
}

struct AccessConfigStruct {
	1:i32 id,
	2:string name,
	3:string source,
	4:i32 addtime
}

struct AccessConfigResponseStruct {	
	1:i32 status,
	2:string msg
	3:AccessConfigStruct accessConfig
}

service AccessConfigThriftService {
    AccessConfigResponseStruct getAccessConfig (1:AccessConfigRequestStruct requestObj)
}

//select * from jl_access_config where source=$source limit 1;
