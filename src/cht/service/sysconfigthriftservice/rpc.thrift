namespace go sysconfigthriftservice 

struct SysConfigStruct {
    1:i32 id,
    2:string nid,
    3:string value,
    4:string name,
}

struct SysConfigRequestStruct {
	1:string chengHuiTongTraceLog
}

struct SysConfigResponseStruct {
    1:i32 status
    2:string msg 
    3:list<SysConfigStruct> sysConfigList
}

service SysConfigThriftService {
    SysConfigResponseStruct getSysConfig(1:SysConfigRequestStruct  requestObj)
}